package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type itemShopRepositoryMock struct {
	mock.Mock
}

func NewItemShopRepositoryMock() *itemShopRepositoryMock {
	return &itemShopRepositoryMock{}
}

func (m *itemShopRepositoryMock) TransactionBegin() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *itemShopRepositoryMock) TransactionRollback(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *itemShopRepositoryMock) TransactionCommit(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *itemShopRepositoryMock) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	args := m.Called(itemFilter)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *itemShopRepositoryMock) FindById(itemID *uint64) (*entities.Item, error) {
	args := m.Called(itemID)
	return args.Get(0).(*entities.Item), args.Error(1)
}

func (m *itemShopRepositoryMock) FindByIdList(itemIDs []*uint64) ([]*entities.Item, error) {
	args := m.Called(itemIDs)
	return args.Get(0).([]*entities.Item), args.Error(1)
}

func (m *itemShopRepositoryMock) PurchaseHistoryRecording(tx *gorm.DB, purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error) {
	args := m.Called(tx, purchasingEntity)
	return args.Get(0).(*entities.PurchaseHistory), args.Error(1)
}