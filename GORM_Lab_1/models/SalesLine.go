package models

type SalesLine struct {
	ID             uint64  `json:"id" gorm:"primaryKey"`
	SalesHeaderID  uint64  `json:"sales_header_id" gorm:"comment:Sipariş başlık ID"`
	ProductID      uint64  `json:"product_id" gorm:"comment:Ürün ID"`
	ProductCode    string  `json:"product_code" gorm:"comment:Ürün Kodu"`
	ProductName    string  `json:"product_name" gorm:"comment:Ürün Adı"`
	ProductBrand   string  `json:"product_brand" gorm:"comment: Ürün Markası"`
	ProductOrigin  string  `json:"product_origin" gorm:"comment: Ürün Menşei"`
	UnitPrice      float64 `json:"unit_price" gorm:"comment:Ürün Fiyatı"`
	Quantity       float64 `json:"quantity" gorm:"comment:Miktar"`
	UnitOfMeasure  string  `json:"unit_of_measure" gorm:"comment:Birim: Hangi birimden satılıyor."`
	DiscountAmount float64 `json:"discount_amount" gorm:"comment:Indirim Tutarı"`
	LineAMount     float64 `json:"line_a_mount" gorm:"comment:Mal hizmet tutarı"`
	VATRate        float64 `json:"vat_rate" gorm:"comment:KDV Oranı"`
	VATAmount      float64 `json:"vat_amount" gorm:"comment:KDV tutarı"`
	Amount         float64 `json:"amount" gorm:"comment:KDV Dahil Tutar"`
}
