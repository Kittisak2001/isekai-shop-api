package services

import (
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/model"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
)

type (
	ItemManagingService interface {
		Creating(itemCreatingModel *model.ItemCreatingReq) (*_itemShopModel.Item, error)
		Editing(itemID *uint64, itemEditingModel *model.ItemEditingReq) (*_itemShopModel.Item, error)
		Archiving(itemID *uint64) error
	}
)