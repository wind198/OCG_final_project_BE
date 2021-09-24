package model

import (
	"OCG_final_project_BE/api/config"
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
	Quantity          int  `json:"quantity"`
}

type OrderReport struct {
	TotalOrders  int     `json:"total_orders"`
	TotalSales   float64 `json:"total_sales"`
	PaidOrders   int     `json:"paid_orders"`
	UnpaidOrders int     `json:"unpaid_orders"`
}

type Duration struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

//  First  i validate, ensure everything from body is valid :D
//  Then calculate total price
//  And we have a new order ting ting! I will wait until my king pay(s) their bill(s)

func GetOneOrder(id string) (Order, error) {
	var order Order
	if err := config.Database.Where("id = ? ", id).First(&order).Error; err != nil {
		fmt.Println(err.Error())
		return order, errors.New("sorry, we have an unexpected error handing your payment. Please come back later")
	}
	return order, nil
}

// Create order
func Create(r *http.Request) (Order, error) {
	order, err := validateOrder(r)
	if err != nil {
		return order, err
	}
	//calculate total price from db

	config.Database.Omit("ID", "FulfilledAt", "ReportSend").Create(&order)
	return order, err
}

func UpdateAfterPay(id string) (time.Time, error) {
	order := Order{}
	time := time.Now().Format("2006-01-02 15:04:05")
	if config.Database.Debug().Model(&order).
		Where("orders.id = ? and orders.fulfilled_at is null ", id).
		Update("fulfilled_at", time).RowsAffected == 0 {
		return order.FulfilledAt, errors.New("order has been paid before :>")
	}
	return order.FulfilledAt, nil
}

func UpdateOrderAfterSend(tx *gorm.DB, st, et string) error {
	err := tx.Debug().Model(&Order{}).
		Where("orders.created_at between ? and ?", st, et).
		Update("orders.report_send", true).Error
	return err
}

// Return unpaid orders, paid orders, total sales for system
// cause we have to check report send or not
func OrderAnalysis(st, et string) (OrderReport, error) {
	var orp OrderReport
	startTime, Endtime, err := ValidateAnalysisQuery(st, et)
	if err != nil {
		return orp, err
	}
	// query with count and sum
	query := "count(id) as total_orders, sum(case when orders.fulfilled_at is not null then orders.total_price else null end) as total_sales, count(case when orders.fulfilled_at is not null then orders.id else null end) as paid_orders, count(case when orders.fulfilled_at is null then orders.id else null end) as unpaid_orders"
	config.Database.Debug().Model(&Order{}).
		Select(query).Where("created_at between ? and ? and report_send is null", startTime, Endtime).Find(&orp)
	return orp, err
}

// Order report for client
func OrderAnalysisClient(st, et string) (OrderReport, error) {
	var orp OrderReport
	startTime, Endtime, err := ValidateAnalysisQuery(st, et)
	if err != nil {
		return orp, err
	}
	// query with count and sum
	query := "count(id) as total_orders, sum(case when orders.fulfilled_at is not null then orders.total_price else null end) as total_sales, count(case when orders.fulfilled_at is not null then orders.id else null end) as paid_orders, count(case when orders.fulfilled_at is null then orders.id else null end) as unpaid_orders"
	config.Database.Debug().Model(&Order{}).
		Select(query).Where("created_at between ? and ?", startTime, Endtime).Find(&orp)
	return orp, err
}

func ValidateAnalysisQuery(st, et string) (time.Time, time.Time, error) {
	// Try to convert string to time format, if err is not nill since its invalid
	const layout = "2006-01-02 15:04:05"
	starttime, errS := time.Parse(layout, st)
	endtime, errE := time.Parse(layout, et)
	if errS != nil {
		return starttime, endtime, errors.New("format start time is wrong")
	}
	if errE != nil {
		return starttime, endtime, errors.New("format end time is wrong")
	}
	if !starttime.Before(endtime) {
		return starttime, endtime, errors.New("start end must before end time")
	}
	return starttime, endtime, nil
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
