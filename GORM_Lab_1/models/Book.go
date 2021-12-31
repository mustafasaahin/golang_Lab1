package models

type Book struct {
	ID     uint64 `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"comment:Kitabın Adı"`
	Author string `json:"author" gorm:"comment:Kitabın Yazarı"`
	Mail   string `json:"mail" gorm:"comment:Yazarın email adresi"`
	Phone  string `json:"phone" gorm:"comment:Yazarın telefon numarası"`
}
