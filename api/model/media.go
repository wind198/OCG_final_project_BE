package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ProductID uint   `json:"ProductID"`
	Image     string `json:"Image"`
}
