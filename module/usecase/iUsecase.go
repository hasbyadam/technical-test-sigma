package usecase

import (
	"github.com/hasbyadam/technical-test-sigma/entity"
	"github.com/hasbyadam/technical-test-sigma/module/store"
)

type Methods struct {
	Stores store.StoreInterface
	Config *entity.Config
}

func New(stores store.StoreInterface, config *entity.Config) UsecaseInterface {
	return &Methods{
		Stores: stores,
		Config: config,
	}
}

type UsecaseInterface interface {
}
