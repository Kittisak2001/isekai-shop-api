package config

import (
	"sync"
	"golang.org/x/oauth2"
)

var (
	once           sync.Once
	configInstance *Config

	PlayerGoogleOAuth2     *oauth2.Config
	AdminGoogleOAuth2      *oauth2.Config
	AccessTokenCookieName  = "act"
	RefreshTokenCookieName = "rft"
	StateCookieName        = "state"

	Letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)