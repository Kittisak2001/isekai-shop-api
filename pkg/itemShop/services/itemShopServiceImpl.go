package services

import _itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}