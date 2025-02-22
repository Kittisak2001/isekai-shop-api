package services

import (
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_playerCoinModel "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
	Buying(buyingReq *_itemShopModel.BuyingReq) (*_playerCoinModel.PlayerCoin, error)
	Selling(sellingReq *_itemShopModel.SellingReq) (*_playerCoinModel.PlayerCoin, error)
}