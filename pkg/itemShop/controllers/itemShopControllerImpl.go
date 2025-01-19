package controllers

import (
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}