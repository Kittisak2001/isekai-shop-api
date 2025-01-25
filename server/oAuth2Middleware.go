package server

import (
	_adminRepository "github.com/Kittisak2001/isekai-shop-api/pkg/admin/repositories"
	_oauth2Middleware "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/middleware"
	_oAuth2Service "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/services"
	_playerRepository "github.com/Kittisak2001/isekai-shop-api/pkg/player/repositories"
)

func (s *echoServer) initOAuth2Middleware() _oauth2Middleware.OAuth2Middleware {
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db)
	oAuth2Service := _oAuth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository, s.app.Logger)
	return _oauth2Middleware.NewOAuth2MiddlewareImpl(s.conf.OAuth2, oAuth2Service)
}