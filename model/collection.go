package model

import (
	"backend/config"
	"fmt"

	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Image          string     `json:"image"`
	CollectionName string     `json:"collection_name"`
	PageID         uint       `json:"page_id"`
	Categories     []Category `gorm:"foreignKey:CollectionID"`
}

var clt Collection

func OneCollection(id string) (Collection, error) {
	if err := config.Database.Where("id = ? ", id).First(&clt).Error; err != nil {
		fmt.Println(err.Error())
	}
	config.Database.Model(&clt).Association("Categories").Find(&clt.Categories)
	return clt, nil
}

// return collection and its categories and all product
func OneCollections(id string) (Collection, error) {
	if err := config.Database.Where("id = ? ", id).First(&clt).Error; err != nil {
		fmt.Println(err.Error())
	}
	err := config.Database.Preload("Categories").Preload("Products").Preload("Image").Preload("ProductVariance").Find(&clt).Error
	return clt, err
}

func OneCollectionCategories(id string) (Collection, error) {
	config.Database.Joins("Categories").Joins("Products").Joins("Images").Joins("ProductVariances").First(&clt, 1)
	return clt, nil
}

func OneCollectionTest(id string) (Collection, error) {
	config.Database.Joins("Product").Find(&clt)
	return clt
}
