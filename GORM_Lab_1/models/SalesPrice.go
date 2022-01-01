package models

import (
	"time"
)

type SalesPrice struct {
	ID              uint64    `json:"id" gorm:"primaryKey"`
	CustomerID      uint64    `json:"customer_id" gorm:"comment: Müşteri id"`
	CustomerNo      string    `json:"customer_no" gorm:"comment:Müşteri No"`
	CustomerName    string    `json:"customer_name" gorm:"comment:Müşterinin adı"`
	ProductID       uint64    `json:"product_id" gorm:"comment: Ürün numarası"`
	ProductCode     string    `json:"product_code" gorm:"comment:Ürün Kodu"`
	ProductName     string    `json:"product_name" gorm:"comment:Ürün adı"`
	UnitPrice       float64   `json:"unit_price" gorm:"comment: Ürün Fiyati"`
	CurrencyCode    string    `json:"currency_code" gorm:"comment:Döviz Kodu: USD,TL,EUR gibi"`
	MinimumQuantity float64   `json:"minimum_quantity" gorm:"comment:Minimum alış miktarı"`
	UnitOfMeasure   string    `json:"unit_of_measure" gorm:"comment:Birim"`
	StartDate       time.Time `json:"start_date" gorm:"comment:Fiyatın geçerlilik başlangıç tarihi: Bu tarihten itibaren bu fiyat geçerli olacak."`
	EndDate         time.Time `json:"finish_date" gorm:"comment:Fiyatın geçerlilik bitiş tarihi. Bu tarihe kadar bu fiyat geçerli olacak."`
}
