package server

import (
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/middleware"
	_playerCoinController "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/controllers"
	_playerCoinRepository "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/repositories"
	_playerCoinService "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/services"
)

func (s *echoServer) initPlayerCoinRouter(m middleware.OAuth2Middleware) {
	playerCoinRepository := _playerCoinRepository.NewPlayerCoinRepositoryImpl(s.db)
	playerCoinService := _playerCoinService.NewPlayerCoinServiceImpl(playerCoinRepository, s.app.Logger)
	playerCoinController := _playerCoinController.NewPlayerCoinController(playerCoinService)

	r := s.app.Group("/v1/player-coins")
	r.POST("", playerCoinController.CoinAdding, m.PlayerGoogleAuthorizing)
	r.GET("", playerCoinController.Showing, m.PlayerGoogleAuthorizing)
}