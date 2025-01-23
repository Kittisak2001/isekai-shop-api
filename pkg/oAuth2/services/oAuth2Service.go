package services

import (
	"context"
	_oAuth2Model "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/model"
	"golang.org/x/oauth2"
)

type OAuth2Service interface {
	PlayerAccountCreating(userInfo *_oAuth2Model.UserInfo) error
	AdminAccountCreating(userInfo *_oAuth2Model.UserInfo) error
	IsThisGuyIsReallyPlayer(playerID string) bool
	IsThisGuyIsReallyAdmin(adminID string) bool
	Callback(ctx context.Context, googleOAuth2 *oauth2.Config, code string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, googleOAuth2 *oauth2.Config, token *oauth2.Token, userInfoUrl string) (*_oAuth2Model.UserInfo, error)
}