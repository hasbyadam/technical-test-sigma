package store

import (
	"context"
	"database/sql"

	"github.com/hasbyadam/technical-test-sigma/connection"
	"github.com/hasbyadam/technical-test-sigma/entity"
)

type Client struct {
	mysql *sql.DB
}

func New(config *entity.Config) StoreInterface {
	return &Client{
		mysql: connection.MySql(config.Database.Mysql),
	}
}

type StoreInterface interface {
	Register(ctx context.Context, req entity.UserDetails) (err error)
	GetUserDetails(ctx context.Context, email string) (res entity.UserDetails, err error)
}
