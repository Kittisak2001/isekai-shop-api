package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type inventoryRepositoryMock struct {
	mock.Mock
}

func NewInventoryRepositoryMock() *inventoryRepositoryMock {
	return &inventoryRepositoryMock{}
}

func (m *inventoryRepositoryMock) Filling(tx *gorm.DB, playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	args := m.Called(tx, playerID, itemID, limit)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}
func (m *inventoryRepositoryMock) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {
	args := m.Called(tx, playerID, itemID, limit)
	return args.Error(0)
}
func (m *inventoryRepositoryMock) PlayerItemCounting(playerID string, itemID uint64) *int64 {
	args := m.Called(playerID, itemID)
	return args.Get(0).(*int64)
}

func (m *inventoryRepositoryMock) Listing(playerID string) ([]*entities.Inventory, error) {
	args := m.Called(playerID)
	return args.Get(0).([]*entities.Inventory), args.Error(1)
}