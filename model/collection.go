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

func OneCollectionCategories(id string) (Collection, error) {
	if err := config.Database.Where("id = ? ", id).First(&clt).Error; err != nil {
		fmt.Println(err.Error())
	}
	config.Database.Model(&clt).Association("Categories").Find(&clt.Categories)
	return clt, nil
}
