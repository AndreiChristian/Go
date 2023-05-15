package main

import (
	"github.com/andreichristian/booksApp/food/api/handlers"
	"github.com/andreichristian/booksApp/food/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.InitDB()

	handlers.RegisterFoodHandlers(r, db)

	r.Run()
}
