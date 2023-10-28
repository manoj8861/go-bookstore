package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() error {
	d, err := gorm.Open("mysql", "manoj:manoj8861@tcp(localhost)/restsql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}
