package seed

import (
	"backend/api/config"
	"backend/api/model"
	"fmt"
)

func CreateTables() {
	// config.Database.Migrator().DropTable(&Order{}, &User{}, &Image{}, &Product{}, &ProductVariance{}, &OrderDetail{}, &User{}, &Page{}, &Collection{}, &Category{}, &CategoryProduct{})
	if err := config.Database.Migrator().AutoMigrate(&model.Order{},
		&model.Product{}, &model.Image{}, &model.ProductVariance{},
		&model.OrderDetail{}, &model.User{}, &model.Page{}, &model.Collection{},
		&model.Category{}); err != nil {
		fmt.Print("Error create  table")
	}
}
