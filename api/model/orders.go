package model

import (
	"backend/api/config"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice   float64       `json:"total_price"`
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
	OrderID           uint `json:"order_id"`
	ProductVarianceID uint `json:"product_variance_id"`
	Quantity          int  `json:"quantity,string"`
}

type OrderReport struct {
	TotalOrders  uint    `json:"total_orders"`
	TotalSales   float64 `json:"total_sales"`
	PaidOrders   uint    `json:"paid_orders"`
	UnpaidOrders uint    `json:"unpaid_orders"`
}

type Duration struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

//  First  i validate, ensure everything from body is valid :D
//  Then calculate total price
//  And we have a new order ting ting! I will wait until my king pay(s) their bill(s)

func Create(r *http.Request) (Order, error) {
	order, err := validateOrder(r)
	if err != nil {
		return order, err
	}
	//calculate total price from db

	config.Database.Omit("ID", "FulfilledAt", "ReportSend").Create(&order)
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
	// Check err for each and cusutomer err for more clearly
	switch {
	case order.Email == "":
		return order, errors.New("email is empty, check it")
	case order.Phone == "":
		return order, errors.New("phone is empty, check it")
	case order.CustomerName == "":
		return order, errors.New("customer_name is empty, check it")
	case order.Address == "":
		return order, errors.New("address is empty, check it")
	case order.TotalPrice != 0:
		return order, errors.New("total_price must empty, check it")
	}
	quantityExceedErrorText := ""
	for _, v := range order.OrderDetails {
		if v.ProductVarianceID == 0 || v.Quantity == 0 {
			return order, errors.New("need to input product_variance_id and quanntily number")
		}
		if v.OrderID != 0 || order.ID != 0 {
			return order, errors.New("can not input order_id, make it naked and i will do it for u :D")
		}
		productVarianceID := strconv.Itoa(int(v.ProductVarianceID))
		productVariance, err := OneProductVariance(productVarianceID)
		if err != nil {
			log.Println("error during querying productvariance, ", err)
		}
		if v.Quantity > productVariance.Inventory {
			product, err := OneProduct(strconv.Itoa(int(productVariance.ProductID)))
			if err != nil {
				log.Println(err)
			}
			quantityExceedErrorText += fmt.Sprintf("Sorry, the quantity of product %q exceeding our inventory, maximum quantity allowed is %v\n", product.Name, productVariance.Inventory)
		}
		productVariancePrice := productVariance.Price
		qf := float64(v.Quantity)
		order.TotalPrice += productVariancePrice * qf
	}
	if quantityExceedErrorText != "" {
		return order, errors.New(quantityExceedErrorText)
	}
	return order, err
}

// Return unpaid orders, paid orders, total sales
func OrderAnalysis(r *http.Request) (OrderReport, error) {
	var orp OrderReport
	dt, err := ValidateAnalysisQuery(r)
	if err != nil {
		return orp, err
	}
	// query with count and sum
	query := "count(id) as total_orders, sum(case when orders.fulfilled_at is not null then orders.total_price else null end) as total_sales, count(case when orders.fulfilled_at is not null then orders.id else null end) as paid_orders, count(case when orders.fulfilled_at is null then orders.id else null end) as unpaid_orders"
	config.Database.Debug().Model(&Order{}).
		Select(query).Where("created_at between ? and ?", dt.StartTime, dt.EndTime).Find(&orp)
	return orp, err
}

func ValidateAnalysisQuery(r *http.Request) (Duration, error) {
	dt := Duration{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dt)
	if err != nil {
		return dt, err
	}
	// Try to convert string to time format, if err is not nill since its invalid
	const layout = "2006-01-02 03:04:05.999"
	_, error := time.Parse(layout, dt.StartTime)

	if error != nil {
		return dt, errors.New("format start time is wrong")
	}
	_, error = time.Parse(layout, dt.EndTime)

	if error != nil {
		return dt, errors.New("format end time is wrong")
	}
	return dt, nil
}
