package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"UserName"`
	PassWord string `json:"PassWord"`
}
