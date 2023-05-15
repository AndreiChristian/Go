package models

import (
	"github.com/jinzhu/gorm"
)

type Food struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
