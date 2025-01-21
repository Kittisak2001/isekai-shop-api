package server

import (
	_itemShopController "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/controllers"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
)

func (s *echoServer) initItemShopRouter() {
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository, s.app.Logger)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router := s.app.Group("/v1/item-shops")
	router.GET("", itemShopController.Listing)
}