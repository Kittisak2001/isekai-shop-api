package services

import (
	"context"
	"encoding/json"
	"io"
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
	playerID := userInfo.ID
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
	adminID := userInfo.ID
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

func (s *googleOAuth2Service) Callback(ctx context.Context, googleOAuth2 *oauth2.Config, code string) (*oauth2.Token, error) {
	token, err := googleOAuth2.Exchange(ctx, code)
	if err != nil {
		s.logger.Errorf("Failed to player exchange code: %s", err.Error())
		return nil, &exception.Unauthorized{}
	}
	return token, nil
}

func (s *googleOAuth2Service) GetUserInfo(ctx context.Context, googleOAuth2 *oauth2.Config, token *oauth2.Token, userInfoUrl string) (*_oAuth2Model.UserInfo, error) {
	client := googleOAuth2.Client(ctx, token)
	resp, err := client.Get(userInfoUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	userInfo := new(_oAuth2Model.UserInfo)
	if err := json.Unmarshal(userInfoBytes, userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (s *googleOAuth2Service) RefreshToken(ctx context.Context, googleOAuth2 *oauth2.Config, token *oauth2.Token) (*oauth2.Token, error) {
	updateToken, err := googleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &exception.Unauthorized{}
	}
	return updateToken, nil
}

func (s *googleOAuth2Service) IsThisGuyIsReallyPlayer(playerID string) bool {
	playerEntity, err := s.playerRepository.FindByID(playerID)
	if err != nil {
		return err == nil
	}
	return playerEntity.ID == playerID
}
func (s *googleOAuth2Service) IsThisGuyIsReallyAdmin(adminID string) bool {
	adminEntity, err := s.adminRepository.FindByID(adminID)
	if err != nil {
		return err == nil
	}
	return adminEntity.ID == adminID
}