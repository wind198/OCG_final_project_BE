package model

import (
	"OCG_final_project_BE/api/config"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string            `json:"Name"`
	Description      string            `json:"Description"`
	ProductVariances []ProductVariance `gorm:"foreignKey:ProductID"`
	Images           []Image           `gorm:"foreignKey:ProductID"`
	Categories       []Category        `gorm:"many2many:category_products;"`
}
type ProductVariance struct {
	gorm.Model
	ProductID   uint    `json:"ProductID"`
	Color       string  `json:"Color"`
	Size        string  `json:"Size"`
	Price       float64 `json:"Price"`
	Inventory   int     `json:"Inventory"`
	OrderDetail OrderDetail
}

type ProductReport struct {
	ID         uint    `json:"ID"`
	Name       string  `json:"Name"`
	AmountSold float64 `json:"AmountSold"`
}

func AllProducts() ([]Product, error) {
	var products []Product

	err := config.Database.Preload("Images").Preload("ProductVariances").Find(&products).Error
	return products, err
}

func OneProduct(id string) (Product, error) {
	var product Product
	if err := config.Database.Where("id = ? ", id).First(&product).Error; err != nil {
		fmt.Println(err.Error())
	}
	err := config.Database.Preload("Images").Preload("ProductVariances").Preload("Categories").Find(&product).Error
	return product, err
}

// Return products based on categories.id
func ProductsBasedCategories(id string, limitNum int, offsNum int) ([]interface{}, error) {
	var products []Product
	productsCount := make([]interface{}, 0)
	if err := config.Database.Debug().Preload("Images").
		Preload("ProductVariances").
		Joins("JOIN category_products on category_products.product_id=products.id").
		Joins("JOIN categories on category_products.category_id=categories.id").
		Where("category_products.category_id=?", id).Limit(limitNum).Offset(offsNum).
		Find(&products).Error; err != nil {
		return productsCount, err
	}
	row := config.Database.Table("products").
		Joins("JOIN category_products on category_products.product_id=products.id").
		Joins("JOIN categories on category_products.category_id=categories.id").
		Where("category_products.category_id=?", id).
		Select("count(products.id) as total").Row()
	var total int
	row.Scan(&total)
	productsCount = append(productsCount, total, products)
	return productsCount, nil
}

// Return products based on collections.id
func ProductsBasedCollection(id string, limitNum int, offsNum int) ([]Product, error) {
	var products []Product
	if err := config.Database.Debug().Preload("Images").
		Preload("ProductVariances").
		Joins("JOIN category_products on category_products.product_id=products.id").
		Joins("JOIN categories on category_products.category_id=categories.id").
		Joins("JOIN collections on collections.id=categories.collection_id").
		Where("collections.id=?", id).Limit(limitNum).Offset(offsNum).
		Find(&products).Error; err != nil {
		log.Fatal(err)
	}
	return products, nil
}

func BestSellProducts(st, et string) ([]ProductReport, error) {
	var products []ProductReport

	startTime, endTime, err := ValidateAnalysisQuery(st, et)
	if err != nil {
		return products, err
	}

	rows, err := config.Database.Table("products").
		Select("products.id,products.name, sum(quantity) total").
		Joins("JOIN order_details on order_details.product_variance_id=products.id").
		Joins("JOIN orders on orders.id=order_details.order_id").
		Group("products.id").Order("total desc").Limit(10).
		Where("order_details.created_at between ? and ? and orders.fulfilled_at is not null", startTime, endTime).Rows()

	if err != nil {
		panic(err)
	}
	var id uint
	var total float64
	var name string
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name, &total)
		if err != nil {
			// handle this error
			panic(err)
		}
		products = append(products, ProductReport{
			ID:         id,
			AmountSold: total,
			Name:       name,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return products, err
}

func OneProductVariance(id string) (ProductVariance, error) {
	var productVariance ProductVariance
	var err error
	if err = config.Database.Where("id = ? ", id).First(&productVariance).Error; err != nil {
		fmt.Println(err.Error())
	}
	return productVariance, err
}
