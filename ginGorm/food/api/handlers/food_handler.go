package handlers

import (
	"net/http"

	"github.com/andreichristian/booksApp/food/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFoodHandlers(r *gin.Engine, db *gorm.DB) {
	r.POST("/foods", func(c *gin.Context) {
		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&food)

		c.JSON(200, food)
	})

	// Add other endpoints here
}
