package models

type Product struct {
	ID                uint64  `json:"id" gorm:"primaryKey"`
	Code              string  `json:"code" gorm:"comment:Ürün kodu"`
	Name              string  `json:"name"gorm:"comment:Ürün adı"`
	Brand             string  `json:"brand" gorm:"comment: Marka"`
	ItemCategory      string  `json:"item_category" gorm:"comment:Ürün kategorisi"`
	ItemSubCategory   string  `json:"item_sub_category" gorm:"comment:Ürün alt kategorisi"`
	Origin            string  `json:"origin" gorm:"comment:Menşei"`
	Weight            float64 `json:"weight" gorm:"comment: Ürünün ağırlıgı"`
	BaseUnitOfMeasure string  `json:"base_unit_of_measure" gorm:"comment:Ana satış birimi"`
	Url               string  `json:"url" gorm:"comment:Ürün web adresi"`
	PictureUrl        string  `json:"picture_url" gorm:"comment: Resim Adresi"`
	UnitPrice         float64 `json:"unit_price" gorm:"comment:Ürün Fiyatı"`
	CurrencyCode      string  `json:"currency_code" gorm:"comment:Döviz Kodu: Örneğin USD,EUR yazılabilir."`
}
