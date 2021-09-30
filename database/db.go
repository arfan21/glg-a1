package database

import (
	"fmt"
	"os"

	"github.com/arfan21/hacktiv8-golang-jwt/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() error {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbname, dbPort)

	gormDB, err := gorm.Open(postgres.Open(config))
	if err != nil {
		return err
	}
	db = gormDB

	db.Debug().AutoMigrate(model.User{}, model.Product{})

	return nil
}

func GetDB() *gorm.DB {
	return db
}
