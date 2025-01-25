package controllers

import "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/services"

type inventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(inventoryService services.InventoryService) *inventoryController {
	return &inventoryController{inventoryService:inventoryService}
}