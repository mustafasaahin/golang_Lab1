package models

type Car struct {
	ID     uint64 `json:"id" gorm:"comment:primaryKey"`
	Mark   string `json:"mark" gorm:"comment:arabanın markası"`
	Year   string `json:"year" gorm:"comment:uretim yili"`
	Nation string `json:"nation"gorm:"comment:Hangi ulkede uretildi"`
}
