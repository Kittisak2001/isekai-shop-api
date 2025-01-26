package tests

import (
	"testing"

	"github.com/Kittisak2001/isekai-shop-api/entities"
	_inventoryRepository "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
	_playerCoinModel "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
	_playerCoinRepository "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/repositories"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestItemBuyingSuccess(t *testing.T) {
	// arrange
	itemShopRepositoryMock := _itemShopRepository.NewItemShopRepositoryMock()
	playerCoinRepositoryMock := _playerCoinRepository.NewPlayerCoinRepositoryMock()
	inventoryRepositoryMock := _inventoryRepository.NewInventoryRepositoryMock()
	logger := echo.New().Logger

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
		logger,
	)

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("TransactionBegin").Return(tx)
	itemShopRepositoryMock.On("TransactionRollback", tx).Return(nil)
	itemShopRepositoryMock.On("TransactionCommit", tx).Return(nil)

	itemID := uint64(1)
	itemShopRepositoryMock.On("FindById", &itemID).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Coin:     5000,
	}, nil)

	itemShopRepositoryMock.On("PurchaseHistoryRecording", tx, &entities.PurchaseHistory{
		PlayerID:        "P001",
		ItemID:          1,
		ItemName:        "Sword of Tester",
		ItemDescription: "A sword that can be used to test the enemy's defense",
		ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
		ItemPrice:       1000,
		Quantity:        3,
		IsBuying:        true,
	}).Return(
		&entities.PurchaseHistory{
			PlayerID:        "P001",
			ItemID:          1,
			ItemName:        "Sword of Tester",
			ItemDescription: "A sword that can be used to test the enemy's defense",
			ItemPicture:     "https://www.google.com/sword-of-tester.jpg",
			ItemPrice:       1000,
			Quantity:        3,
			IsBuying:        true,
		}, nil)

	inventoryRepositoryMock.On("Filling", tx, "P001", uint64(1), int(3)).Return([]*entities.Inventory{
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
		{
			PlayerID: "P001",
			ItemID:   1,
		},
	}, nil)

	playerCoinRepositoryMock.On("CoinAdding", tx, &entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}).Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Amount:   -3000,
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected *_playerCoinModel.PlayerCoin
	}

	// act
	cases := []args{
		{
			label: "Test Item Buying Success",
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &_playerCoinModel.PlayerCoin{
				PlayerID: "P001",
				Amount:   -3000,
			},
		},
	}

	// assert
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.NoError(t, err)
			assert.EqualValues(t, c.expected, result)

		})
	}
}

func TestItemBuyingFail(t *testing.T) {
	// arrange
	itemShopRepositoryMock := _itemShopRepository.NewItemShopRepositoryMock()
	playerCoinRepositoryMock := _playerCoinRepository.NewPlayerCoinRepositoryMock()
	inventoryRepositoryMock := _inventoryRepository.NewInventoryRepositoryMock()
	logger := echo.New().Logger

	itemShopService := _itemShopService.NewItemShopServiceImpl(
		itemShopRepositoryMock,
		playerCoinRepositoryMock,
		inventoryRepositoryMock,
		logger,
	)

	tx := &gorm.DB{}
	itemShopRepositoryMock.On("TransactionBegin").Return(tx)
	itemShopRepositoryMock.On("TransactionRollback", tx).Return(nil)
	itemShopRepositoryMock.On("TransactionCommit", tx).Return(nil)
	itemID := uint64(1)
	itemShopRepositoryMock.On("FindById", &itemID).Return(&entities.Item{
		ID:          1,
		Name:        "Sword of Tester",
		Price:       1000,
		Description: "A sword that can be used to test the enemy's defense",
		Picture:     "https://www.google.com/sword-of-tester.jpg",
	}, nil)

	playerCoinRepositoryMock.On("Showing", "P001").Return(&entities.PlayerCoin{
		PlayerID: "P001",
		Coin:     2000,
	}, nil)

	type args struct {
		label    string
		in       *_itemShopModel.BuyingReq
		expected error
	}

	// act
	cases := []args{
		{
			label: "CoinNotEnough",
			in: &_itemShopModel.BuyingReq{
				PlayerID: "P001",
				ItemID:   1,
				Quantity: 3,
			},
			expected: &exception.CoinNotEnough{},
		},
	}

	// assert
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			result, err := itemShopService.Buying(c.in)
			assert.Nil(t, result)
			assert.Error(t, err)
			assert.EqualValues(t, c.expected, err)
		})
	}
}