package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type inventoryRepositoryImpl struct {
	db databases.Database
}

func NewInventoryRepositoryImpl(db databases.Database) InventoryRepository {
	return &inventoryRepositoryImpl{db: db}
}

func (r *inventoryRepositoryImpl) Filling(tx *gorm.DB, playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	inventoryEntities := make([]*entities.Inventory, 0)
	for range limit {
		inventoryEntities = append(inventoryEntities, &entities.Inventory{
			PlayerID: playerID,
			ItemID:   itemID,
		})
	}
	if err := conn.CreateInBatches(inventoryEntities, len(inventoryEntities)).Error; err != nil {
		return nil, err
	}
	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) Removing(tx *gorm.DB, playerID string, itemID uint64, limit int) error {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}
	inventoryEntities, err := r.findPlayerItemInInventoryByID(playerID, itemID, limit)
	if err != nil {
		return err
	}
	for _, inventory := range inventoryEntities {
		inventory.IsDeleted = true
		if err := conn.Model(&entities.Inventory{}).Where("id = ?", inventory.ID).Updates(inventory).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *inventoryRepositoryImpl) findPlayerItemInInventoryByID(playerID string, itemID uint64, limit int) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)
	if err := r.db.Connect().Where("player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false).Limit(limit).Find(&inventoryEntities).Error; err != nil {
		return nil, err
	}
	return inventoryEntities, nil
}

func (r *inventoryRepositoryImpl) PlayerItemCounting(playerID string, itemID uint64) *int64 {
	count := new(int64)
	if err := r.db.Connect().Model(&entities.Inventory{}).Where("player_id = ? AND item_id = ? AND is_deleted = ?", playerID, itemID, false).Count(count).Error; err != nil {
		return nil
	}
	return count
}

func (r *inventoryRepositoryImpl) Listing(playerID string) ([]*entities.Inventory, error) {
	inventoryEntities := make([]*entities.Inventory, 0)
	if err := r.db.Connect().Where("player_id = ? AND is_deleted = ?", playerID, false).Find(&inventoryEntities).Error; err != nil {
		return nil, err
	}
	return inventoryEntities, nil
}
