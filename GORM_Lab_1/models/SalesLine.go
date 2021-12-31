package models

type SalesLine struct {
	ID               uint64  `json:"id" gorm:"primaryKey"`
	ProductID        uint64  `json:"product_id" gorm:"comment:Ürün ID"`
	ProductCode      string  `json:"product_code" gorm:"comment:Ürün Kodu"`
	ProductName      string  `json:"product_name" gorm:"comment:Ürün Adı"`
	ProductBrand     string  `json:"product_brand" gorm:"comment: Ürün Markası"`
	ProductOrigin    string  `json:"product_origin" gorm:"comment: Ürün Menşei"`
	UnitPrice        float64 `json:"unit_price" gorm:"comment:Ürün Fiyatı"`
	Quantity         float64 `json:"quantity" gorm:"comment:Miktar"`
	DiscountQuantity float64 `json:"discount_quantity" gorm:"comment:Indirim Tutarı"`
	VATRate          float64 `json:"vat_rate" gorm:"comment:KDV Oranı"`
	VATQuantity      float64 `json:"vat_quantity" gorm:"comment:KDV miktarı"`
	LineAMount       float64 `json:"line_a_mount" gorm:"comment:Mal hizmet tutarı"`
}
