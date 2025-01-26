package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	"gorm.io/gorm"
)

type ItemShopRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	FindById(itemID *uint64) (*entities.Item, error)
	FindByIdList(itemIDs []*uint64) ([]*entities.Item, error)
	PurchaseHistoryRecording(tx *gorm.DB,purchasingEntity *entities.PurchaseHistory) (*entities.PurchaseHistory, error)
	Counting(itemFilter *_itemShopModel.ItemFilter) (*int64, error)
}