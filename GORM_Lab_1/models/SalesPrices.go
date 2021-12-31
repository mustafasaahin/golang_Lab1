package models

import (
	"time"
)

type SalesPrices struct {
	ID              uint64    `json:"id" gorm:"primaryKey"`
	CustomerID      uint64    `json:"customer_id" gorm:"comment: Müşteri no"`
	ProductID       uint64    `json:"product_id" gorm:"comment: Ürün numarası"`
	UnitPrice       float64   `json:"unit_price" gorm:"comment: Ürün Fiyati"`
	MinimumQuantity float64   `json:"minimum_quantity" gorm:"comment:Minimum alış miktarı"`
	UnitOfMeasure   string    `json:"unit_of_measure" gorm:"comment:Birim"`
	StartDate       time.Time `json:"start_date" gorm:"comment:Fiyat başlangıç tarihi"`
	FinishDate      time.Time `json:"finish_date" gorm:"comment:Fiyat bitiş tarihi"`
}
