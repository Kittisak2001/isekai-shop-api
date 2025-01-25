package repositories

type InventoryRepository interface{
	FillingItem()
	FindItemByName()
	RemovingItem()
}