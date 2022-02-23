package models

// Dev : Mustafa ŞAHİN
// Date 23.02.2022
// Description : musteriler

type Customer struct {
	ID                 uint64    `json:"id" gorm:"primaryKey"`
	No                 string    `json:"no" gorm:"comment:Müşteri No"`
	Name               string    `json:"name" gorm:"comment:Müşterinin adi"`
	ContactName        string    `json:"contact_name" gorm:"comment:İlgili kişi"`
	Phone              string    `json:"phone" gorm:"comment:Müşterinin telefonu"`
	Fax                string    `json:"fax" gorm:"comment:Fax no"`
	Mail               string    `json:"mail" gorm:"comment:Email adresi"`
	PaymentTerm        uint8     `json:"payment_term" gorm:"comment:Ödeme vadesi"`
	CurrencyCode       string    `json:"currency_code" gorm:"comment:Döviz cinsi"`
	Address            string    `json:"adress" gorm:"comment:Müşteri adresi"`
	City               string    `json:"city" gorm:"comment:Müşteri Şehri"`
	Country            string    `json:"country" gorm:"comment:Müşteri ülkesi"`
	TaxAreaCode        string    `json:"tax_area_code" gorm:"comment: Vergi Dairesi"`
	TaxNo              string    `json:"tax_no" gorm:"comment: Vergi No"`
	TradeRegisterNo    string    `json:"trade_register_no" gorm:"comment:Ticaret sicil no"`
	Type               string    `json:"type" gorm:"comment:Müşteri tipi"`
	PersonalIdentityNo string    `json:"personal_identity_no" gorm:"comment:Müşteri TC no"`
	WebAddress         string    `json:"web_address" gorm:"comment:Web Adresi"`
	Iban               string    `json:"iban" gorm:"comment: Müşteri Iban"`
	Contacts           []Contact `json:"contacts" gorm:"foreignKey:customer_id;comment:Müşterinin ilgili kişi listesi"`
}
