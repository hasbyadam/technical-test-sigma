package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // sql driver
	"github.com/hasbyadam/technical-test-sigma/entity"
	"go.uber.org/zap"
)

// MySql
func MySql(opts entity.Mysql) *sql.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", opts.User, opts.Password, opts.Host, opts.Dbname)

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		zap.S().DPanic(err)
	}

	err = db.Ping()
	if err != nil {
		zap.S().DPanic(err)
	}

	zap.S().Info("Connected to MySql DB Server")

	return db
}
