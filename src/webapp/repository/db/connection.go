package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// NewDbconnection returns a db connection object
func NewDbconnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Andres82463.@tcp(127.0.0.1:3306)/db1")
	CheckErr(err)
	err = db.Ping()
	CheckErr(err)
	return db
}

//CheckErr function print the error on console
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
