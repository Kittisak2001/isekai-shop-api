package repositories

import (
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db *gorm.DB
}

func NewItemShopRepositoryImpl(db *gorm.DB) ItemShopRepository {
	return &itemShopRepositoryImpl{db}
}

func (r *itemShopRepositoryImpl) Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // SELECT * FROM items
	name := itemFilter.Name
	if name != ""{
		query = query.Where("name like ?", "%"+name+"%") // WHERE name LIKE %..%
	}
	description := itemFilter.Description
	if description != ""{
		query = query.Where("description like ?", "%"+description+"%") // WHERE description LIKE %..%
	}
	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)
	if err := query.Offset(offset).Limit(limit).Order("id asc").Find(&itemList).Error; err != nil {
		return nil, err
	}
	return itemList, nil
}

func (r *itemShopRepositoryImpl) Counting(itemFilter *_itemShopModel.ItemFilter) (*int64, error) {
	query := r.db.Model(&entities.Item{}).Where("is_archive = ?", false) // SELECT * FROM items
	name := itemFilter.Name
	if name != ""{
		query = query.Where("name like ?", "%"+name+"%") // WHERE name LIKE %..%
	}
	description := itemFilter.Description
	if description != ""{
		query = query.Where("description like ?", "%"+description+"%") // WHERE description LIKE %..%
	}
	count := new(int64)
	if err := query.Count(count).Error; err != nil {
		return nil, err
	}
	return count, nil
}