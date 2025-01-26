package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Filling(tx *gorm.DB, playerID string, itemID uint64, limit int) ([]*entities.Inventory, error)
	Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error
	PlayerItemCounting(playerID string, itemID uint64) *int64
	Listing(playerID string) ([]*entities.Inventory, error)
}