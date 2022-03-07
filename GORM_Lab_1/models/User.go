package models

import (
	"time"
)

type User struct {
	ID                 uint64    `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name" gorm:"Comment:Kullanici ismi"`
	Surname            string    `json:"surname" gorm:"Comment:Kullanici soyadi"`
	BirthDate          time.Time `json:"birth_date" gorm:"Comment:Kullanici dogum tarihi"`
	CompanyName        string    `json:"company_name" gorm:"Comment:Şirket adi"`
	Mail               string    `json:"mail" gorm:"Comment:Kullanici mail adresi"`
	Phone              string    `json:"phone" gorm:"Comment: Kullanici telefon numarasi"`
	PersonalIdentityNo string    `json:"personal_identity_no" gorm:"Comment: Kullanici TC kimlik numarasi"`
	Gender             string    `json:"gender" gorm:"Comment:Cinsiyet"`
	Password           string    `json:"password" gorm:"Comment:Parola"`
}
