package services

import (
	"context"
	_oAuth2Model "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/model"
	"golang.org/x/oauth2"
)

type OAuth2Service interface {
	PlayerAccountCreating(userInfo *_oAuth2Model.UserInfo) error
	AdminAccountCreating(userInfo *_oAuth2Model.UserInfo) error
	Callback(ctx context.Context, googleOAuth2 *oauth2.Config, code string, userInfoUrl string) (*oauth2.Token, *_oAuth2Model.UserInfo, error)
}