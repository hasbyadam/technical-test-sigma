package main

import (
	// "time"

	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hasbyadam/technical-test-sigma/entity"
	appInit "github.com/hasbyadam/technical-test-sigma/init"
	"github.com/hasbyadam/technical-test-sigma/module/handler"
	"github.com/hasbyadam/technical-test-sigma/module/store"
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	config *entity.Config
)

func init() {
	appInit.SetupLogger()
	config = appInit.SetupMainConfig()
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger(), middleware.RequestID(), middleware.Secure())
	main := e.Group("/")

	
	handler.New(main, &usecase.Methods{
		Stores: store.New(config),
		Config: config,
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(config.API.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), config.Context.Timeout*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
