package usecase

import (
	"context"

	"github.com/hasbyadam/technical-test-sigma/entity"
	"github.com/hasbyadam/technical-test-sigma/module/store"
	"github.com/hasbyadam/technical-test-sigma/schema/request"
	"github.com/hasbyadam/technical-test-sigma/schema/response"
)

type Methods struct {
	Stores store.StoreInterface
	Config *entity.Config
	Claims entity.Claims
}

func New(stores store.StoreInterface, config *entity.Config) UsecaseInterface {
	return &Methods{
		Stores: stores,
		Config: config,
	}
}

type UsecaseInterface interface {
	Register(ctx context.Context, req request.Register) (res response.Register, err error)
	Login(ctx context.Context, req request.Login) (res response.Login, err error)
	GetConfig() *entity.Config
	ParseTokenToClaims(tokenString string) (err error)
	RequestTransaction(ctx context.Context, req request.RequestTransaction) (res response.RequestTransaction, err error)
}
