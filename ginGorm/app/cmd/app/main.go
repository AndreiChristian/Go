package main

import (
	"github.com/andreichristian/ginGorm/app/api/handlers"
	"github.com/andreichristian/ginGorm/app/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := db.InitDB()

	handlers.RegisterFoodHandlers(r, db)

	r.Run()
}
