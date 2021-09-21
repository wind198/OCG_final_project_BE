package model

import (
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
	ProductVarianceID uint    `json:"product_id"`
	Quantity          float32 `json:"quantity"`
}
