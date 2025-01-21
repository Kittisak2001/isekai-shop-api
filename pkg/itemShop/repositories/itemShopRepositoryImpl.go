package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
)

type itemShopRepositoryImpl struct {
	db databases.Database
}

func NewItemShopRepositoryImpl(db databases.Database) ItemShopRepository {
	return &itemShopRepositoryImpl{db}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)
	query := r.db.Connect().Model(&entities.Item{}).Where("is_archive = ?", false) // SELECT * FROM items
	name := itemFilter.Name
	if name != "" {
		query = query.Where("name like ?", "%"+name+"%") // WHERE name LIKE %..%
	}
	description := itemFilter.Description
	if description != "" {
		query = query.Where("description like ?", "%"+description+"%") // WHERE description LIKE %..%
	}
	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)
	if err := query.Offset(offset).Limit(limit).Order("id asc").Find(&itemList).Error; err != nil {
		return nil, err
	}
	return itemList, nil
}

func (r *itemShopRepositoryImpl) FindById(itemID *uint64) (*entities.Item, error) {
	itemEntity := new(entities.Item)
	if err := r.db.Connect().Where("id = ?", itemID).Take(itemEntity).Error; err != nil {
		return nil, err
	}
	return itemEntity, nil
}
