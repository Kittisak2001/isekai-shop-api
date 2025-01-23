package middleware

import "github.com/labstack/echo/v4"

type OAuth2Middleware interface {
	PlayerGoogleAuthorizing(next echo.HandlerFunc) echo.HandlerFunc
	AdminGoogleAuthorizing(next echo.HandlerFunc) echo.HandlerFunc
}