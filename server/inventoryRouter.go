package server

import (
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/controllers"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/repositories"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/services"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/middleware"
)

func (s *echoServer) initInventoryRouter(m middleware.OAuth2Middleware) {
	r := s.app.Group("/v1/inventory")

	inventoryRepository := repositories.NewInventoryRepositoryImpl(s.db)
	itemShopRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db)
	inventoryService := services.NewInventoryServiceImpl(inventoryRepository, itemShopRepository,s.app.Logger)
	inventoryController := controllers.NewInventoryController(inventoryService)
	r.GET("", inventoryController.Listing, m.PlayerGoogleAuthorizing)
}