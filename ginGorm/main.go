package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=containers-us-west-27.railway.app user=postgres password=0jEOklCsjnt1s3OuPh6n dbname=postgres port=6876 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connected")

}
