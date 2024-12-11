package handler

import (
	"context"

	"github.com/hasbyadam/technical-test-sigma/pkg/util"
	"github.com/hasbyadam/technical-test-sigma/schema/request"
	"github.com/hasbyadam/technical-test-sigma/schema/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(c echo.Context) error {
	req := request.Register{}
	res := response.Register{}

	err := c.Bind(&req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	res, err = h.Usecase.Register(context.Background(), req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success register", res)
}

func (h *Handler) Login(c echo.Context) error {
	req := request.Login{}
	res := response.Login{}

	err := c.Bind(&req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	res, err = h.Usecase.Login(context.Background(), req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success login", res)
}
