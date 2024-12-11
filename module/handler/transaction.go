package handler

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hasbyadam/technical-test-sigma/pkg/util"
	"github.com/hasbyadam/technical-test-sigma/schema/request"
	"github.com/hasbyadam/technical-test-sigma/schema/response"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RequestTransaction(c echo.Context) error {
	req := request.RequestTransaction{}
	res := response.RequestTransaction{}

	err := c.Bind(&req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	if err = h.Usecase.ParseTokenToClaims(c.Get("user").(*jwt.Token).Raw); err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}
	res, err = h.Usecase.RequestTransaction(context.Background(), req)
	if err != nil {
		return util.ErrorInternalServerResponse(c, err, res)
	}

	return util.SuccessResponse(c, "success register", res)
}
