package server

import (
	_itemManagingController "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/controllers"
	_itemManagingRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/repositories"
	_itemManagingService "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/services"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
)

func (s *echoServer) initItemManagingRouter() {
	r := s.app.Group("/v1/item-managing")
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db)
	itemManagingRepository := _itemManagingRepository.NewItemManagingRepository(s.db)
	itemManagingService := _itemManagingService.NewItemManagingService(itemManagingRepository, itemShopRepository, s.app.Logger)
	itemManagingController := _itemManagingController.NewItemManagingController(itemManagingService)
	r.POST("", itemManagingController.Creating)
	r.PATCH("/:itemID", itemManagingController.Editing)
	r.DELETE("/:itemID", itemManagingController.Archiving)
}
