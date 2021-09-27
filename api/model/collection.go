package model

import (
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Image          string     `json:"Image"`
	CollectionName string     `json:"CollectionName"`
	PageID         uint       `json:"PageID"`
	Categories     []Category `gorm:"foreignKey:CollectionID"`
}

// func OneCollectionCategories(id string) (Collection, error) {
// 	var clt Collection
// 	err := config.Database.Debug().Preload("Categories").Where("ID=?", id).First(&clt).Error
// 	return clt, err
// }
