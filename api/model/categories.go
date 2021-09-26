package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"CategoryName"`
	CollectionID uint      `json:"CollectionID"`
	Products     []Product `gorm:"many2many:category_products;"`
}
