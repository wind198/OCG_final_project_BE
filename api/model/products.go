package model

import (
	"backend/api/config"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ProductVariances []ProductVariance `gorm:"foreignKey:ProductID"`
	Images           []Image           `gorm:"foreignKey:ProductID"`
	Categories       []Category        `gorm:"many2many:category_products;"`
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

var product Product

func AllProducts() ([]Product, error) {
	var products []Product

	err := config.Database.Preload("Images").Preload("ProductVariances").Find(&products).Error
	return products, err
}

func OneProduct(id string) (Product, error) {
	if err := config.Database.Where("id = ? ", id).First(&product).Error; err != nil {
		fmt.Println(err.Error())
	}
	err := config.Database.Preload("Images").Preload("ProductVariances").Preload("Categories").Find(&product).Error
	return product, err
}

func ProductsBasedCategories(id string, limitNum int, offsNum int) ([]Product, error) {
	var products []Product
	if err := config.Database.Debug().Preload("Images").
		Preload("ProductVariances").
		Joins("JOIN category_products on category_products.product_id=products.id").
		Joins("JOIN categories on category_products.category_id=categories.id").
		Where("category_products.category_id=?", id).Limit(limitNum).Offset(offsNum).
		Find(&products).Error; err != nil {
		log.Fatal(err)
	}
	return products, nil
}
