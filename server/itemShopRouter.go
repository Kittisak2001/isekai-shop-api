package server

import (
	_inventoryRepository "github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	_itemShopController "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/controllers"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/middleware"
	_playerCoinRepository "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/repositories"
)

func (s *echoServer) initItemShopRouter(m middleware.OAuth2Middleware) {
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db)
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db)
	inventoryRepository := _inventoryRepository.NewInventoryRepositoryImpl(s.db)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository, playerCoinRepository, inventoryRepository, s.app.Logger)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router := s.app.Group("/v1/item-shops")
	router.GET("", itemShopController.Listing)
	router.POST("/buying", itemShopController.Buying, m.PlayerGoogleAuthorizing)
	router.POST("/selling", itemShopController.Selling, m.PlayerGoogleAuthorizing)
}