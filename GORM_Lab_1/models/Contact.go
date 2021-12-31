package models

type Contact struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	CustomerID   uint64 `json:"customer_id" gorm:"comment:Müşteri ID"`
	CustomerName string `json:"customer_name" gorm:"comment:Müşteri Adı"`
	FirstName    string `json:"first_name" gorm:"comment:İlgili kişiş adı"`
	LastName     string `json:"last_name" gorm:"comment:İlgili kişi soyadı"`
	Phone        string `json:"phone" gorm:"comment:Telefon numarası"`
	MobilPhone   string `json:"mobil_phone" gorm:"comment:Cep telefon numarası"`
	Mail         string `json:"mail" gorm:"comment:Mail adresi"`
	Gender       string `json:"gender" gorm:"comment:Cinseti"`
}
