package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "127.0.0.1", "3306", "test")

	gormDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	db = gormDB
	return nil
}

func GetDB() *gorm.DB {
	return db
}
