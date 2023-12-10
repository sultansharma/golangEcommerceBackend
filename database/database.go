package database

import (
	"flutter-dev.info/go-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbase, err := gorm.Open(mysql.Open("meru:123456@/meru-bazar?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("DB error")
	}
	DB = dbase
	dbase.AutoMigrate(&models.User{})
	dbase.AutoMigrate(&models.Shop{})
	dbase.AutoMigrate(&models.Product{})
	dbase.AutoMigrate(&models.ProductVariation{})
}
