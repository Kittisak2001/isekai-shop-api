package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kittisak2001/isekai-shop-api/entities"
	_adminRepository "github.com/Kittisak2001/isekai-shop-api/pkg/admin/repositories"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/exception"
	_oAuth2Model "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/model"
	_playerRepository "github.com/Kittisak2001/isekai-shop-api/pkg/player/repositories"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type (
	googleOAuth2Service struct {
		playerRepository _playerRepository.PlayerRepository
		adminRepository  _adminRepository.AdminRepository
		logger           echo.Logger
	}
)

func NewGoogleOAuth2Service(playerRepository _playerRepository.PlayerRepository, adminRepository _adminRepository.AdminRepository, logger echo.Logger) OAuth2Service {
	return &googleOAuth2Service{playerRepository, adminRepository, logger}
}

func (s *googleOAuth2Service) PlayerAccountCreating(userInfo *_oAuth2Model.UserInfo) error {
	playerID := userInfo.Sub
	if !s.isThisGuyIsPlayerReally(playerID) {
		playerEntity := &entities.Player{
			ID:     playerID,
			Email:  userInfo.Email,
			Name:   userInfo.Name,
			Avatar: userInfo.Picture,
		}
		_, err := s.playerRepository.Creating(playerEntity)
		if err != nil {
			s.logger.Errorf("Failed to creating player: %s", err.Error())
			return &exception.PlayerCreating{PlayerID: playerID}
		}
	}
	return nil
}

func (s *googleOAuth2Service) isThisGuyIsPlayerReally(playerID string) bool {
	player, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		s.logger.Errorf("Failed to FindByID: %s", err.Error())
		return err == nil
	}
	return player != nil
}

func (s *googleOAuth2Service) AdminAccountCreating(userInfo *_oAuth2Model.UserInfo) error {
	adminID := userInfo.Sub
	if !s.isThisGuyIsAdminReally(adminID) {
		adminEntity := &entities.Admin{
			ID:     adminID,
			Email:  userInfo.Email,
			Name:   userInfo.Name,
			Avatar: userInfo.Picture,
		}
		_, err := s.adminRepository.Creating(adminEntity)
		if err != nil {
			s.logger.Errorf("Failed to creating player: %s", err.Error())
			return &exception.AdminCreating{AdminID: adminID}
		}
	}
	return nil
}

func (s *googleOAuth2Service) isThisGuyIsAdminReally(adminID string) bool {
	admin, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return err == nil
	}
	return admin != nil
}

func (s *googleOAuth2Service) Callback(ctx context.Context, googleOAuth2 *oauth2.Config, code string, userInfoUrl string) (*oauth2.Token, *_oAuth2Model.UserInfo, error) {
	token, err := googleOAuth2.Exchange(ctx, code)
	if err != nil {
		s.logger.Errorf("Failed to player exchange code: %s", err.Error())
		return nil, nil, &exception.Unauthorized{}
	}
	client := googleOAuth2.Client(ctx, token)
	userInfo, err := s.getUserInfo(client, userInfoUrl)
	if err != nil {
		s.logger.Errorf("Error getting user info: %s", err.Error())
		return nil, nil, &exception.Unauthorized{}
	}
	return token, userInfo, nil
}

func (s *googleOAuth2Service) getUserInfo(client *http.Client, userInfoUrl string) (*_oAuth2Model.UserInfo, error) {
	resp, err := client.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	s.logger.Info("Body =>", string(userInfoBytes))
	userInfo := new(_oAuth2Model.UserInfo)
	if err := json.Unmarshal(userInfoBytes, userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
