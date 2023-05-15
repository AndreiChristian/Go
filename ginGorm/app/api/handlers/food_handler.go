package handlers

import (
	"net/http"

	"github.com/andreichristian/ginGorm/app/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterFoodHandlers(r *gin.Engine, db *gorm.DB) {

	r.GET("/foods", func(c *gin.Context) {
		var foods []models.Food
		db.Find(&foods)
		c.JSON(200, foods)
	})

	r.POST("/foods", func(c *gin.Context) {
		var food models.Food
		if err := c.ShouldBindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Create(&food)

		c.JSON(200, food)
	})

	r.DELETE("/foods/:id", func(c *gin.Context) {
		var food models.Food
		id := c.Param("id")

		db.First(&food, id)

		if food.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No food found!"})
			return
		}

		db.Delete(&food)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Food deleted successfully!"})
	})

}
