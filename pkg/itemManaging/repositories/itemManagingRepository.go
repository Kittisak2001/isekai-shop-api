package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemManagingModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/model"
)

type (
	ItemManagingRepository interface {
		Creating(itemEntity *entities.Item) error
		Editing(itemID *uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (error)
		Archiving(itemID *uint64) error
	}
)