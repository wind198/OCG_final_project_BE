package model

import (
	"backend/api/config"

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
