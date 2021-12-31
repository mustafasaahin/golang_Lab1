package models

import "time"

type SalesHeader struct {
	ID                         uint64    `json:"id" gorm:"primaryKey"`
	CustomerID                 uint64    `json:"customer_id" gorm:"comment: Müşteri ID"`
	InvoiceNo                  string    `json:"invoice_no" gorm:"comment:Fatura No"`
	PostingDate                time.Time `json:"posting_date" gorm:"comment: fatura tarihi`
	OrderDate                  time.Time `json:"order_date"gorm:"comment:Sipariş tarihi"`
	CustomerPersonalIdentityNo string    `json:"customer_personal_identity_no" gorm:"comment:Müşteri TC no"`
	CustomerName               string    `json:"customer_name" gorm:"comment: Müşteri Adı"`
	CustomerContactName        string    `json:"customer_contact_name" gorm:"comment:İlgili kişi adı`
	CustomerPhone              string    `json:"customer_phone" gorm:"comment:Müşteri Telefonu"`
	CustomerFax                string    `json:"customer_fax" gorm:"comment:Müşteri fax no"`
	TaxAreCode                 string    `json:"tax_are_code"gorm:"comment:Vergi Dairesi"`
	TaxNo                      string    `json:"tax_no"gorm:"comment:Vergi No"`
	CustomerMail               string    `json:"customer_mail" gorm:"comment:Müşteri maili"`
	PaymentTerm                uint8     `json:"payment_term" gorm:"comment: ödeme vadesi"`
	CustomerAddress            string    `json:"customer_address" gorm:"comment: Müşteri adresi"`
	CustomerCity               string    `json:"customer_city" gorm:"comment:Müşteri şehri"`
	CustomerCountry            string    `json:"customer_country" gorm:"comment: Müşteri ülkesi"`
	TotalAmount                float64   `json:"total_amount" gorm:"comment: Toplam tutar"`
	TotalVAT                   float64   `json:"total_vat" gorm:"commment: KDV tutarı"`
	PaymentMethod              string    `json:"payment_method" gorm:"comment: Ödeme şekli"`
	DueDate                    time.Time `json:"due_date" gorm:"comment:Ödeme tarihi"`
}
