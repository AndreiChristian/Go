package db

import (
	"github.com/andreichristian/ginGorm/app/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=containers-us-west-27.railway.app user=postgres dbname=railway password=0jEOklCsjnt1s3OuPh6n sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Food{})
	return db
}
