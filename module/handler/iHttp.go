package handler

import (
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

func New(e *echo.Group, usecase usecase.UsecaseInterface) {
	handler := &Handler{
		Usecase: usecase,
	}
	
	v1 := e.Group("/v1")
	
	auth := v1.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	trans := v1.Group("/transaction")
	trans.Use(echojwt.JWT([]byte(handler.Usecase.GetConfig().Auth.JwtSecretKey)))
	
	trans.POST("", handler.RequestTransaction)
}
