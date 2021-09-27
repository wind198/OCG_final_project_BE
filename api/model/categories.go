package model

import (
	"OCG_final_project_BE/api/config"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"CategoryName"`
	CollectionID uint      `json:"CollectionID"`
	Products     []Product `gorm:"many2many:category_products;"`
}

func OneCollectionCategories(id string) ([]Category, error) {
	var ctg []Category
	if err := config.Database.Debug().
		Joins("JOIN category_products on category_products.product_id=products.id").
		Joins("JOIN categories on category_products.category_id=categories.id").
		Joins("JOIN collections on collections.id=categories.collection_id").
		Where("categories.collection_id=?", id).
		Find(&ctg).Error; err != nil {
		return ctg, err
	}
	return ctg, nil
}
