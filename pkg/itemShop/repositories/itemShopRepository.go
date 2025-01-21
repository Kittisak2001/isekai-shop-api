package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	FindById(itemID *uint64) (*entities.Item, error)
}