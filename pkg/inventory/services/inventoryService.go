package services

import "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/model"

type InventoryService interface {
	Listing(playerID string) ([]*model.Inventory, error)
}