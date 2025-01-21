package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
	db   databases.Database
}

var (
	once   sync.Once
	server *echoServer
)

func NewEchoServer(conf *config.Config, db databases.Database) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	once.Do(func() {
		server = &echoServer{
			app:  echoApp,
			conf: conf,
			db:   db,
		}
	})
	return server
}

func (s *echoServer) Start() {
	corsMiddleware := getCORSMiddleware(s.conf.Server.AllowOrigins)
	timeOutMiddleware := getTimeOutMiddleware(s.conf.Server.Timeout)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(middleware.BodyLimit(s.conf.Server.BodyLimit))
	s.app.Use(timeOutMiddleware)

	// health check ไว้ใช้กับ cloud provider
	s.app.GET("/v1/health", s.healthCheck)
	s.app.GET("/v1/panic", func(c echo.Context) error {
		panic("panic")
	})

	s.initItemShopRouter()
	s.initItemManagingRouter()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyShutdown(quitCh)
	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) gracefullyShutdown(quitCh chan os.Signal) {
	// Waiting
	<-quitCh
	s.app.Logger.Info("Shutting down server...")
	if err := s.app.Shutdown(context.Background()); err != nil {
		s.app.Logger.Fatalf("Error: %s", err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func getTimeOutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(
		middleware.TimeoutConfig{
			Skipper:      middleware.DefaultSkipper,
			ErrorMessage: "Request Timeout",
			Timeout:      timeout * time.Second,
		},
	)
}

func getCORSMiddleware(allowOrigins []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(
		middleware.CORSConfig{
			Skipper:      middleware.DefaultSkipper,
			AllowOrigins: allowOrigins,
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		},
	)
}
