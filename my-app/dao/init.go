package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/gin_study?charset=utf8mb4"

	//db, err := sql.Open("mysql", dsn)

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	log.Println("DB connect successfully")

	DB = db

}
