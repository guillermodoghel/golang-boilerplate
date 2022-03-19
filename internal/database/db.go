package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	mysqlUrlFmt    = "%s:%s@(%s:%s)/%s?parseTime=true&charset=utf8mb4"
	mysqlUrlFmtRaw = "%s:%s@(%s:%s)/%s?charset=utf8mb4"
	maxIdleConns   = 20
	maxOpenConns   = 150
)

func GetDB() (*sqlx.DB, error) {
	var (
		dbUser     = "root"
		dbPassword = "root"
		dbHost     = "localhost"
		dbPort     = "3306"
		dbName     = "iq"
	)
	databaseUrl := fmt.Sprintf(mysqlUrlFmt, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	return db, nil
}