package database

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	err_db error
)

func DbConnect() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mysql",
		Addr:                 "localhost",
		DBName:               "stockwave",
		AllowNativePasswords: true,
	}

	db, err_db = sql.Open("mysql", cfg.FormatDSN())
	if err_db != nil {
		log.Fatal(err_db)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Printf("Conectado a %s.", cfg.DBName)
}

func GetConnectDb() *sql.DB {
	return db
}
