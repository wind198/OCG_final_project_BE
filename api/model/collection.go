package model

import (
	"OCG_final_project_BE/api/config"

	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Image          string     `json:"image"`
	CollectionName string     `json:"collection_name"`
	PageID         uint       `json:"page_id"`
	Categories     []Category `gorm:"foreignKey:CollectionID"`
}

func OneCollectionCategories(id string) (Collection, error) {
	var clt Collection
	err := config.Database.Debug().Preload("Categories").Where("ID=?", id).First(&clt).Error
	return clt, err
}

type CollectionCreate struct {
	Colections Collection
	Categories []Category
}

func OneCollectionProduct(id string) ([]Collection, error) {
	var collections []Collection
	err := config.Database.Where("id = ?", id).
		Preload("Categories.Products.ProductVariances").
		Preload("Categories.Products.Images").
		First(&collections).
		Error
	return collections, err
}
