package store

import (
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
	
}
