package repositories

import "gorm.io/gorm"

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) FillingItem() {
	panic("not yet")
}
func (r *inventoryRepository) FindItemByName() {
	panic("not yet")
}
func (r *inventoryRepository) RemovingItem() {
	panic("not yet")
}