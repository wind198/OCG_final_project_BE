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
