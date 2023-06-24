package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"my-app/logs"
)

var DB *gorm.DB

func InitDB() {
	dsn := "TheCSA:123456@tcp(127.0.0.1:3306)/thecsa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.LogError(err)
	}
	if err != nil {
		logs.LogError(err)
	}

	log.Println("DB connect successfully")

	DB = db

}
