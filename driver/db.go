package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	err error
)

// Init is initialize db from main function
func Init() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	db, err = sqlx.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/"+dbName+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// GetDB is called in models
func GetDB() *sqlx.DB {
	return db
}

// Close DB connection
func Close() {
	db.Close()
}
