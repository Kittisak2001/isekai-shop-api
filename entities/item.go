package entities

import (
	"time"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
)

type (
	Item struct {
		ID          uint64    `gorm:"primaryKey;autoIncrement;"`
		AdminID     *string   `gorm:"type:varchar(64);"`
		Name        string    `gorm:"type:varchar(64);unique;not null;"`
		Description string    `gorm:"type:varchar(128);not null;"`
		Price       uint       `gorm:"not null;default:0;"`
		Picture     string    `gorm:"not null;"`
		IsArchive   bool      `gorm:"not null;default:false;"`
		CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
		UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
	}
)

func (i *Item) ToItemModel() *_itemShopModel.Item {
	return &_itemShopModel.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Price:       i.Price,
		Picture:     i.Picture,
	}
}