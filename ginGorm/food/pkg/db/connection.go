package db

import (
	"github.com/andreichristian/booksApp/food/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	dsn := "host=containers-us-west-27.railway.app user=postgres password=0jEOklCsjnt1s3OuPh6n dbname=postgres port=6876 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Food{})

	return db

}
