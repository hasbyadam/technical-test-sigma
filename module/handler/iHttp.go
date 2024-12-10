package handler

import (
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
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

	v1.POST("/register", handler.Register)
	v1.POST("/login", handler.Login)
}
