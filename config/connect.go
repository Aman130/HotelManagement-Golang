package config

import (
	"github.com/Aman130/HotelManagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=hotel password=hotel dbname=hotel port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Hotel{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Bookings{})
	DB = db
}
