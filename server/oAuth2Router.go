package server

import (
	_playerRepository "github.com/Kittisak2001/isekai-shop-api/pkg/player/repositories"
	_adminRepository "github.com/Kittisak2001/isekai-shop-api/pkg/admin/repositories"
	_oAuth2Controller "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/controllers"
	_oAuth2Service "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/services"
)

func (s *echoServer)initOAuth2Router(){
	playerRepository := _playerRepository.NewPlayerRepositoryImpl(s.db)
	adminRepository := _adminRepository.NewAdminRepositoryImpl(s.db)
	oAuth2Service := _oAuth2Service.NewGoogleOAuth2Service(playerRepository, adminRepository, s.app.Logger)
	oAuth2Controller := _oAuth2Controller.NewGoogleOAuth2Controller(oAuth2Service, s.conf.OAuth2, s.app.Logger)
	
	r := s.app.Group("/v1/oauth2/google")
	admin := r.Group("/admin")
	adminLogin := admin.Group("/login")
	adminLogin.GET("", oAuth2Controller.AdminLogin)
	adminLogin.GET("/callback", oAuth2Controller.AdminLoginCallback)
	
	player := r.Group("/player")
	playerLogin := player.Group("/login")
	playerLogin.GET("", oAuth2Controller.PlayerLogin)
	playerLogin.GET("/callback", oAuth2Controller.PlayerLoginCallback)

	r.DELETE("/logout", oAuth2Controller.Logout)
}