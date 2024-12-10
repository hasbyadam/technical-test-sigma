package handler

import (
	"github.com/hasbyadam/technical-test-sigma/module/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase usecase.UsecaseInterface
}

func New(e *echo.Group, usecase usecase.UsecaseInterface) {
	// handler := &Handler{
	// 	Usecase: usecase,
	// }
}
