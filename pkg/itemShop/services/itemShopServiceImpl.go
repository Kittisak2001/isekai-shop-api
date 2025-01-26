package services

import (
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_inventoryException "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/exceptions"
	_inventoryRepository "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	_itemShopException "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	_playercoinException "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/exception"
	_playerCoinModel "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/repositories"
	"github.com/labstack/echo/v4"
)

type itemShopServiceImpl struct {
	itemShopRepository   _itemShopRepository.ItemShopRepository
	playerCoinRepository _playerCoinRepository.PlayerCoinRepository
	inventoryRepository  _inventoryRepository.InventoryRepository
	logger               echo.Logger
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository, playerCoinRepository _playerCoinRepository.PlayerCoinRepository, inventoryRepository _inventoryRepository.InventoryRepository, logger echo.Logger) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository, playerCoinRepository, inventoryRepository, logger}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		s.logger.Errorf("Failed to list item or counting: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}
	itemCounting := int64(len(itemEntityList))
	totalPage := s.totalPageCalculation(&itemCounting, itemFilter.Size)
	return s.toItemResultResponse(itemEntityList, itemFilter.Page, totalPage), err
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems *int64, size int64) int64 {
	totalPage := *totalItems / size
	if *totalItems%size != 0 {
		totalPage++
	}
	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}

// Buying
// -> return player coin
func (s *itemShopServiceImpl) Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error) {
	// -> find item by id *
	itemEntity, err := s.itemShopRepository.FindById(&buyingReq.ItemID)
	if err != nil {
		s.logger.Errorf("Error to find itemID: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{ItemID: buyingReq.ItemID}
	}
	// -> total price calculation *
	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), buyingReq.Quantity)
	// -> check player coin *
	if err := s.playerCoinChecking(buyingReq.PlayerID, totalPrice); err != nil {
		s.logger.Errorf("Error to checking player coin:  %s", err.Error())
		return nil, err
	}
	// -> purchase history *
	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        buyingReq.PlayerID,
		ItemID:          buyingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        buyingReq.Quantity,
		IsBuying:        config.Buying,
	})
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error history of purchase recording: %s", err.Error())
		return nil, &_itemShopException.HistoryOfPurchaseRecording{}
	}
	s.logger.Infof("Purchase histort recored: %s", purchaseRecording.ID)

	// -> coin decuting *
	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: buyingReq.PlayerID,
		Amount:   -totalPrice,
	})
	if err != nil {
		s.logger.Errorf("Player coin decuted: %d", playerCoin.Amount)
		return nil, &_playercoinException.CoinAdding{}
	}
	s.logger.Infof("Player coin decuted: %d", playerCoin.Amount)
	// -> inventory filling
	inventoryEntity, err := s.inventoryRepository.Filling(tx, buyingReq.PlayerID, buyingReq.ItemID, int(buyingReq.Quantity))
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error inventory filling: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{PlayerID: buyingReq.PlayerID, ItemID: buyingReq.ItemID}
	}
	s.logger.Infof("Inventory failed: %d", len(inventoryEntity))
	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error transaction commit: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{}
	}
	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) totalPriceCalculation(item *_itemShopModel.Item, qty uint) int64 {
	return int64(item.Price) * int64(qty)
}

func (s *itemShopServiceImpl) playerCoinChecking(playerID string, totalPrice int64) error {
	playerCoin, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return err
	}
	if playerCoin.Coin < totalPrice {
		return &_itemShopException.CoinNotEnough{}
	}
	return nil
}

// Buying
// -> find item by id -> total price calculation -> check player item -> purchase history -> coin adding -> inventory removing -> return player coin
func (s *itemShopServiceImpl) Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error) {
	// -> find item by id *
	itemEntity, err := s.itemShopRepository.FindById(&sellingReq.ItemID)
	if err != nil {
		s.logger.Errorf("Error to find itemID: %s", err.Error())
		return nil, &_itemShopException.ItemNotFound{ItemID: sellingReq.ItemID}
	}
	// -> total price calculation *
	totalPrice := s.totalPriceCalculation(itemEntity.ToItemModel(), sellingReq.Quantity) / 2
	// -> check player item *
	if err := s.playerItemChecking(sellingReq.PlayerID, sellingReq.ItemID, sellingReq.Quantity); err != nil {
		s.logger.Errorf("Error to player item checking:  %s", err.Error())
		return nil, err
	}
	// -> purchase history *
	tx := s.itemShopRepository.TransactionBegin()
	purchaseRecording, err := s.itemShopRepository.PurchaseHistoryRecording(tx, &entities.PurchaseHistory{
		PlayerID:        sellingReq.PlayerID,
		ItemID:          sellingReq.ItemID,
		ItemName:        itemEntity.Name,
		ItemDescription: itemEntity.Description,
		ItemPrice:       itemEntity.Price,
		ItemPicture:     itemEntity.Picture,
		Quantity:        sellingReq.Quantity,
		IsBuying:        config.Selling,
	})
	if err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error history of purchase recording: %s", err.Error())
		return nil, &_itemShopException.HistoryOfPurchaseRecording{}
	}
	s.logger.Infof("Purchase histort recored: %s", purchaseRecording.ID)

	// -> coin decuting *
	playerCoin, err := s.playerCoinRepository.CoinAdding(tx, &entities.PlayerCoin{
		PlayerID: sellingReq.PlayerID,
		Amount:   totalPrice,
	})
	if err != nil {
		s.logger.Errorf("Player coin decuted: %d", playerCoin.Amount)
		return nil, &_playercoinException.CoinAdding{}
	}
	s.logger.Infof("Player coin decuted: %d", playerCoin.Amount)
	// -> inventory filling
	if err := s.inventoryRepository.Removing(tx, sellingReq.PlayerID, sellingReq.ItemID, int(sellingReq.Quantity)); err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error inventory filling: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{PlayerID: sellingReq.PlayerID, ItemID: sellingReq.ItemID}
	}
	s.logger.Infof("Inventory itemID: %d, removed: %d", sellingReq.ItemID, sellingReq.Quantity)
	if err := s.itemShopRepository.TransactionCommit(tx); err != nil {
		s.itemShopRepository.TransactionRollback(tx)
		s.logger.Errorf("Error transaction commit: %s", err.Error())
		return nil, &_inventoryException.InventoryFilling{}
	}
	return playerCoin.ToPlayerCoinModel(), nil
}

func (s *itemShopServiceImpl) playerItemChecking(playerID string, itemID uint64, qty uint) error {
	itemCounting := s.inventoryRepository.PlayerItemCounting(playerID, itemID)
	if uint(*itemCounting) < qty {
		return &_itemShopException.ItemNotEnough{ItemID: itemID}
	}
	return nil
}