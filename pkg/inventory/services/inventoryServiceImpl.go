package services

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/exceptions"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/model"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	"github.com/labstack/echo/v4"
)

type inventoryServiceImpl struct {
	inventoryRepository repositories.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
	logger              echo.Logger
}

func NewInventoryServiceImpl(inventoryRepository repositories.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository, logger echo.Logger) InventoryService {
	return &inventoryServiceImpl{inventoryRepository: inventoryRepository, itemShopRepository: itemShopRepository, logger: logger}
}

func (s *inventoryServiceImpl) Listing(playerID string) ([]*model.Inventory, error) {
	inventoryEntities, err := s.inventoryRepository.Listing(playerID)
	if err != nil {
		s.logger.Errorf("Error to liting inventory of PlayerID: %s", playerID)
		return nil, &exceptions.InventoryListing{playerID}
	}
	uniqueItemWithQuantityCounterList := s.getUniqueItemWithQuantityCounterList(inventoryEntities)
	return s.buildInventoryListingResult(uniqueItemWithQuantityCounterList), nil
}

func (s *inventoryServiceImpl) getUniqueItemWithQuantityCounterList(inventoryEntities []*entities.Inventory) []*model.ItemQuantityCounting {
	itemQuantityCounterList := make([]*model.ItemQuantityCounting, 0)

	itemMapWithQuantity := make(map[uint64]uint)

	for _, inventory := range inventoryEntities {
		itemMapWithQuantity[inventory.ItemID]++
	}

	for itemID, quantity := range itemMapWithQuantity {
		itemQuantityCounterList = append(itemQuantityCounterList, &model.ItemQuantityCounting{
			ItemID:   itemID,
			Quantity: quantity,
		})
	}
	return itemQuantityCounterList
}

func (s *inventoryServiceImpl) buildInventoryListingResult(uniqueItemWithQuantityCounterList []*model.ItemQuantityCounting) []*model.Inventory {
	uniqueItemIDList := s.getItemByID(uniqueItemWithQuantityCounterList)
	itemEntities, err := s.itemShopRepository.FindByIdList(uniqueItemIDList)
	if err != nil {
		return nil
	}
	results := make([]*model.Inventory, 0)
	itemMapWithQuantity := s.getItemMapWithQuantity(uniqueItemWithQuantityCounterList)
	for _, itemEntity := range itemEntities {
		results = append(results, &model.Inventory{
			Item:     *itemEntity.ToItemModel(),
			Quantity: itemMapWithQuantity[itemEntity.ID],
		})
	}
	return results
}

func (s *inventoryServiceImpl) getItemByID(uniqueItemWithQuantityCounterList []*model.ItemQuantityCounting) []*uint64 {
	uniqueItemIDList := make([]*uint64, 0)
	for _, inventory := range uniqueItemWithQuantityCounterList {
		uniqueItemIDList = append(uniqueItemIDList, &inventory.ItemID)
	}
	return uniqueItemIDList
}

func (s *inventoryServiceImpl) getItemMapWithQuantity(uniqueItemWithQuantityCounterList []*model.ItemQuantityCounting) map[uint64]uint {
	itemMapWithQuantity := make(map[uint64]uint)
	for _, v := range uniqueItemWithQuantityCounterList {
		itemMapWithQuantity[v.ItemID] += v.Quantity
	}
	return itemMapWithQuantity
}