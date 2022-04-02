package dao

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"log"

	//"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var dB *gorm.DB

func InitDB() {
	//db, err := sql.Open("mysql", "root:xianye@tcp(localhost)/db1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	sqlDB, err := sql.Open("mysql", "root:xianye@tcp(localhost)/db1")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dB = gormDB

}
