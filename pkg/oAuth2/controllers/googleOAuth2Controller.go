package controllers

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/model"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
	"golang.org/x/oauth2"
)

type (
	googleOAuth2Controller struct {
		oAuth2Service services.OAuth2Service
		oAuth2Conf    *config.OAuth2Cfg
		logger        echo.Logger
	}
)

var (
	playerGoogleOAuth2 *oauth2.Config
	adminGoogleOAuth2  *oauth2.Config
	once               sync.Once

	accessTokenCookieName  = "act"
	refreshTokenCookieName = "rft"
	stateCookieName        = "state"

	letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func NewGoogleOAuth2Controller(oAuth2Service services.OAuth2Service, oAuth2Conf *config.OAuth2Cfg, logger echo.Logger) *googleOAuth2Controller {
	once.Do(func() {
		setGoogleOAuth2Config(oAuth2Conf)
	})

	return &googleOAuth2Controller{oAuth2Service, oAuth2Conf, logger}
}

func setGoogleOAuth2Config(oauth2Conf *config.OAuth2Cfg) {
	playerGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.PlayerRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}

	adminGoogleOAuth2 = &oauth2.Config{
		ClientID:     oauth2Conf.ClientID,
		ClientSecret: oauth2Conf.ClientSecret,
		RedirectURL:  oauth2Conf.AdminRedirectUrl,
		Scopes:       oauth2Conf.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       oauth2Conf.Endpoints.AuthUrl,
			TokenURL:      oauth2Conf.Endpoints.TokenUrl,
			DeviceAuthURL: oauth2Conf.Endpoints.DeviceAuthUrl,
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func (c *googleOAuth2Controller) PlayerLogin(pctx echo.Context) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	state := c.randomState()
	c.setCookie(pctx, stateCookieName, state)
	return pctx.Redirect(http.StatusFound, playerGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) AdminLogin(pctx echo.Context) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	state := c.randomState()
	c.setCookie(pctx, stateCookieName, state)
	return pctx.Redirect(http.StatusFound, adminGoogleOAuth2.AuthCodeURL(state))
}

func (c *googleOAuth2Controller) PlayerLoginCallback(pctx echo.Context) error {
	if err := c.callbackValidating(pctx); err != nil {
		c.logger.Errorf("Error validating callback: %s", err.Error())
		return custom.Error(pctx, http.StatusUnauthorized, &exception.FailCallback{Role: "player"})
	}

	token, userInfo, err := c.oAuth2Service.Callback(pctx.Request().Context(), playerGoogleOAuth2, pctx.QueryParam("code"), c.oAuth2Conf.UserInfoUrl)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, &exception.FailCallback{Role: "player"})
	}
	if err := c.oAuth2Service.PlayerAccountCreating(userInfo); err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, &exception.PlayerCreating{PlayerID: userInfo.Sub})
	}
	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)
	return pctx.JSON(http.StatusOK, &model.LoginResponse{Message: "Login success"})
}

func (c *googleOAuth2Controller) AdminLoginCallback(pctx echo.Context) error {
	token, userInfo, err := c.oAuth2Service.Callback(pctx.Request().Context(), adminGoogleOAuth2, pctx.QueryParam("code"), c.oAuth2Conf.UserInfoUrl)
	if err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, &exception.FailCallback{Role: "player"})
	}
	if err := c.oAuth2Service.PlayerAccountCreating(userInfo); err != nil {
		return custom.Error(pctx, http.StatusUnauthorized, &exception.PlayerCreating{PlayerID: userInfo.Sub})
	}
	c.setSameSiteCookie(pctx, accessTokenCookieName, token.AccessToken)
	c.setSameSiteCookie(pctx, refreshTokenCookieName, token.RefreshToken)
	return pctx.JSON(http.StatusOK, &model.LoginResponse{Message: "Login success"})
}

func (c *googleOAuth2Controller) callbackValidating(pctx echo.Context) error {
	state := pctx.QueryParam("state")
	stateFromCookie, err := pctx.Cookie(stateCookieName)
	if err != nil {
		return &exception.ProcessCookie{}
	}
	if state != stateFromCookie.Value {
		return &exception.InvalidState{}
	}
	c.removeCookie(pctx, stateCookieName)
	return nil
}

func (c *googleOAuth2Controller) Logout(pctx echo.Context) error {
	accessToken, err := pctx.Cookie(accessTokenCookieName)
	if err != nil && err.Error() != "http: named cookie not present" {
		c.logger.Errorf("Error reading access token: %s", err.Error())
		return &exception.ProcessCookie{}
	}
	if accessToken != nil {
		if err := c.revokeToken(accessToken.Value); err != nil {
			c.logger.Errorf("Error revoking token: %s", err.Error())
			return &exception.ProcessCookie{}
		}
	}
	c.removeSameSiteCookie(pctx, accessTokenCookieName)
	c.removeSameSiteCookie(pctx, refreshTokenCookieName)
	return pctx.JSON(http.StatusOK, &model.LogoutResponse{Message: "Logout successful"})
}

func (c *googleOAuth2Controller) revokeToken(accessToken string) error {
	revokeURL := fmt.Sprintf("%s?token=%s", c.oAuth2Conf.RevokeUrl, accessToken)
	resp, err := http.Post(revokeURL, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return &exception.FailRevoke{}
	}
	defer resp.Body.Close()

	return nil
}

func (c *googleOAuth2Controller) setCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		// Secure: true,
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		// Secure: true,
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) setSameSiteCookie(pctx echo.Context, name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure: true,
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) removeSameSiteCookie(pctx echo.Context, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		// Secure: true,
	}
	pctx.SetCookie(cookie)
}

func (c *googleOAuth2Controller) randomState() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
