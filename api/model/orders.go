package model

import (
	"backend/api/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice   float32       `json:"total_price"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	CustomerName string        `json:"customer_name"`
	Address      string        `json:"address"`
	ReportSend   string        `json:"report_send"`
	FulfilledAt  time.Time     `json:"fulfilled_at"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	gorm.Model
	OrderID           uint    `json:"order_id"`
	ProductVarianceID uint    `json:"product_variance_id"`
	Quantity          float32 `json:"quantity"`
}

func Create(r *http.Request) (Order, error) {
	fmt.Println("Run create order")
	order, err := validateOrder(r)
	config.Database.Omit("ID", "FulfilledAt", "ReportSend").Create(&order)
	// save all fields when performing incluces time...
	return order, err
}

func validateOrder(r *http.Request) (Order, error) {

	decoder := json.NewDecoder(r.Body)
	var order Order
	err := decoder.Decode(&order)
	if err != nil {
		return order, err
	}

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	switch {
	case order.Email == "":
		return order, errors.New("email is empty, check it")
	case order.Phone == "":
		return order, errors.New("phone is empty, check it")
	case order.CustomerName == "":
		return order, errors.New("customer_name is empty, check it")
	case order.Address == "":
		return order, errors.New("address is empty, check it")
	case order.TotalPrice == 0:
		return order, errors.New("total_price is empty, check it")
	}
	for _, v := range order.OrderDetails {
		if v.ProductVarianceID == 0 || v.Quantity == 0 {
			return order, errors.New("need to input product_variance_id and quanntily number")
		}
		if v.OrderID != 0 || order.ID != 0 {
			return order, errors.New("can not input order_id, make it naked and i will do it for u :D")
		}
	}
	return order, nil
}
