package services

import (
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	"github.com/labstack/echo/v4"
)

type inventoryService struct {
	inventoryRepository repositories.InventoryRepository
	itemShopRepository  _itemShopRepository.ItemShopRepository
	logger              echo.Logger
}

func NewInventoryService(inventoryRepository repositories.InventoryRepository, itemShopRepository _itemShopRepository.ItemShopRepository, logger echo.Logger) InventoryService {
	return &inventoryService{inventoryRepository: inventoryRepository, itemShopRepository: itemShopRepository, logger: logger}
}

func (s inventoryService) AddItem() {
	panic("not yet")
}
func (s inventoryService) RemoveItem() {
	panic("not yet")
}