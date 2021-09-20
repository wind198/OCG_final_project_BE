package model

import (
	"backend/config"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ProductVariances []ProductVariance `gorm:"foreignKey:ProductID"`
	Images           []Image           `gorm:"foreignKey:ProductID"`
}
type ProductVariance struct {
	gorm.Model
	ProductID   uint    `json:"product_id"`
	Color       string  `json:"color"`
	Size        string  `json:"size"`
	Price       float32 `json:"price"`
	Inventory   int     `json:"inventory"`
	OrderDetail OrderDetail
}

func AllProducts() ([]Product, error) {
	products := make([]Product, 0)
	err := config.Database.Preload("Images").Preload("ProductVariances").Find(&products).Error
	return products, err
}

var product Product

func OneProduct(id string) (Product, error) {
	if err := config.Database.Where("id = ? ", id).First(&product).Error; err != nil {
		fmt.Println(err.Error())
	}
	err := config.Database.Preload("Images").Preload("ProductVariances").Find(&product).Error
	return product, err
}

var products []Product

func ProductsBasedCategories(id string) ([]Product, error) {
	config.Database.Joins("category_products").Where("category_id = ?", id).Find(&product)
	return products, nil
}
