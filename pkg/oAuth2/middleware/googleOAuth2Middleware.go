package middleware

import (
	"net/http"
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/exception"
	_oAuth2Service "github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type oAuth2MiddlewareImpl struct {
	oAuth2Conf    *config.OAuth2Cfg
	oAuth2Service _oAuth2Service.OAuth2Service
}

func NewOAuth2MiddlewareImpl(oAuth2Conf *config.OAuth2Cfg, oAuth2Service _oAuth2Service.OAuth2Service) OAuth2Middleware {
	return &oAuth2MiddlewareImpl{oAuth2Conf: oAuth2Conf, oAuth2Service: oAuth2Service}
}

func (m *oAuth2MiddlewareImpl) PlayerGoogleAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		ctx := pctx.Request().Context()
		tokenSource, err := m.getTokenSource(pctx)
		if err != nil {
			return custom.Error(pctx, http.StatusInternalServerError, err)
		}
		if !tokenSource.Valid() {
			tokenSource, err = config.PlayerGoogleOAuth2.TokenSource(ctx, tokenSource).Token()
			if err != nil {
				return custom.Error(pctx, http.StatusUnauthorized, err)
			}
			m.setSameSiteCookie(pctx, config.AccessTokenCookieName, tokenSource.AccessToken)
			m.setSameSiteCookie(pctx, config.RefreshTokenCookieName, tokenSource.RefreshToken)
			userInfo, err := m.oAuth2Service.GetUserInfo(ctx, config.PlayerGoogleOAuth2, tokenSource, m.oAuth2Conf.UserInfoUrl)
			if err != nil {
				return custom.Error(pctx, http.StatusUnauthorized, err)
			}
			if !m.oAuth2Service.IsThisGuyIsReallyPlayer(userInfo.ID) {
				return custom.Error(pctx, http.StatusUnauthorized, &exception.Unauthorized{})
			}
			pctx.Set("playerID", userInfo.ID)
		}
		return next(pctx)
	}
}

func (m *oAuth2MiddlewareImpl) AdminGoogleAuthorizing(next echo.HandlerFunc) echo.HandlerFunc {
	return func(pctx echo.Context) error {
		ctx := pctx.Request().Context()
		tokenSource, err := m.getTokenSource(pctx)
		if err != nil {
			return custom.Error(pctx, http.StatusInternalServerError, err)
		}
		if !tokenSource.Valid() {
			tokenSource, err = config.AdminGoogleOAuth2.TokenSource(ctx, tokenSource).Token()
			if err != nil {
				return custom.Error(pctx, http.StatusUnauthorized, err)
			}
			m.setSameSiteCookie(pctx, config.AccessTokenCookieName, tokenSource.AccessToken)
			m.setSameSiteCookie(pctx, config.RefreshTokenCookieName, tokenSource.RefreshToken)
			userInfo, err := m.oAuth2Service.GetUserInfo(ctx, config.PlayerGoogleOAuth2, tokenSource, m.oAuth2Conf.UserInfoUrl)
			if err != nil {
				return custom.Error(pctx, http.StatusUnauthorized, err)
			}
			if !m.oAuth2Service.IsThisGuyIsReallyAdmin(userInfo.ID) {
				return custom.Error(pctx, http.StatusUnauthorized, &exception.Unauthorized{})
			}
			pctx.Set("adminID", userInfo.ID)
		}
		return next(pctx)
	}
}

func (m *oAuth2MiddlewareImpl) RefreshingToken(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	panic("not yet")
}

func (m *oAuth2MiddlewareImpl) adminTokenRefreshing(pctx echo.Context, token *oauth2.Token) (*oauth2.Token, error) {
	ctx := pctx.Request().Context()
	updateToken, err := config.AdminGoogleOAuth2.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, &exception.Unauthorized{}
	}
	m.setSameSiteCookie(pctx, config.AccessTokenCookieName, updateToken.AccessToken)
	m.setSameSiteCookie(pctx, config.RefreshTokenCookieName, updateToken.RefreshToken)
	return updateToken, nil
}

func (m *oAuth2MiddlewareImpl) getTokenSource(pctx echo.Context) (*oauth2.Token, error) {
	accessToken, err := pctx.Cookie(config.AccessTokenCookieName)
	if err != nil {
		return nil, &exception.Unauthorized{}
	}
	refreshToken, err := pctx.Cookie(config.RefreshTokenCookieName)
	if err != nil {
		return nil, &exception.Unauthorized{}
	}
	return &oauth2.Token{
		AccessToken:  accessToken.Value,
		RefreshToken: refreshToken.Value,
	}, nil
}

func (m *oAuth2MiddlewareImpl) setSameSiteCookie(pctx echo.Context, name, value string) {
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

func (m *oAuth2MiddlewareImpl) removeSameSiteCookie(pctx echo.Context, name string) {
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
