package model

import "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"

type (
	Inventory struct {
		Item     model.Item `json:"item"`
		Quantity uint       `json:"quantity"`
	}

	ItemQuantityCounting struct{
		ItemID uint64
		Quantity uint
	}
)