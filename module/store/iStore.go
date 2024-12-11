package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
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
	GetUserLimitBalance(ctx context.Context, userId, creditLimitId uuid.UUID) (res entity.UserLimitBalance, err error)
	InsertTransactions(ctx context.Context, req entity.TransactionDetails) (err error)
}

func LockTables(db *sql.DB, tableName string) error {

	queryString := fmt.Sprintf("lock tables %s write", tableName)

	_, err := db.Exec(queryString)

	if err != nil {
		return err
	}

	return nil
}

func UnlockTables(db *sql.DB) error {

	queryString := "unlock tables"

	_, err := db.Exec(queryString)

	if err != nil {
		return err
	}

	return nil
}
