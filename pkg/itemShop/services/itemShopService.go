package services

import (
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}