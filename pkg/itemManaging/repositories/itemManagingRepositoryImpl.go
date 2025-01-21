package repositories

import (
	"errors"
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemManagingModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/model"
)

type (
	itemManagingRepositoryImpl struct {
		db databases.Database
	}
)

func NewItemManagingRepository(db databases.Database) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) error {
	return r.db.Connect().Create(itemEntity).Error
}

func (r *itemManagingRepositoryImpl) Editing(itemID *uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) error {
	result := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ? AND id = ?", false, itemID).Updates(itemEditingReq)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID *uint64) error {
	return r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error
}
