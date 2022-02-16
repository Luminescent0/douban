package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:xianye@tcp(localhost)/db1")
	if err != nil {
		log.Fatal(err)
	}
	dB = db

}
