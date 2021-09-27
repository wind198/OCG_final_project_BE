package model

import (
	"OCG_final_project_BE/api/config"
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"CategoryName"`
	CollectionID uint      `json:"CollectionID"`
	Products     []Product `gorm:"many2many:category_products;"`
}

type CategoryNumber struct {
	ID            int       `json:"ID"`
	CategoryName  string    `json:"CategoryName"`
	CollectionID  uint      `json:"CollectionID"`
	ProductNumber string    `json:"ProductNumber"`
	Products      []Product `gorm:"many2many:category_products;"`
}

func OneCollectionCategories(id string) ([]CategoryNumber, error) {
	var ctg []Category
	var cn []CategoryNumber
	if err := config.Database.Debug().
		Joins("JOIN collections on collections.id=categories.collection_id").
		Where("categories.collection_id=?", id).
		Find(&ctg).Error; err != nil {
		return cn, err
	}
	for _, v := range ctg {
		fmt.Println(v.CategoryName)
		row := config.Database.Table("products").
			Joins("join category_products on category_products.product_id=products.id").
			Joins("JOIN categories on categories.id=category_products.category_id").
			Joins("JOIN collections on collections.id=categories.collection_id").
			Select("count(products.id) as total").
			Where("categories.id=?", v.ID).Row()
		var number string
		row.Scan(&number)
		cn = append(cn, CategoryNumber{
			ID:            int(v.ID),
			CategoryName:  v.CategoryName,
			CollectionID:  v.CollectionID,
			ProductNumber: number,
		})
	}
	fmt.Print(cn, "\n")
	return cn, nil
}
