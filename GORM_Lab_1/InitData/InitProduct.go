package InitData

import (
	"fmt"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
)

func InitProduct() {
	config.DB.
		Where("id <> 0").
		Delete(&models.Product{})

	product := models.Product{
		Code:              "PLS001",
		Name:              "Plus Banyo Bataryası",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Mix Bataryalar",
		ItemSubCategory:   "Plus Serisi",
		Origin:            "Çin",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/116/plus-banyo-bataryasi",
		PictureUrl:        "http://www.artiarmatur.com/images/product/7e725dbd-e809-.jpg",
		UnitPrice:         245,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

	product = models.Product{
		Code:              "PLS002",
		Name:              "Plus Lavabo Bataryası",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Mix Bataryalar",
		ItemSubCategory:   "Plus Serisi",
		Origin:            "Çin",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/117/plus-lavabo-bataryasi",
		PictureUrl:        "http://www.artiarmatur.com/images/product/07bac228-f0ae-.jpg",
		UnitPrice:         325,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

	product = models.Product{
		Code:              "PLS003",
		Name:              "Plus Kuğu Lavabo Bataryası",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Mix Bataryalar",
		ItemSubCategory:   "Plus Serisi",
		Origin:            "Çin",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/118/plus-kugu-lavabo-bataryasi",
		PictureUrl:        "http://www.artiarmatur.com/images/product/f6041221-5f45-.jpg",
		UnitPrice:         325,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

	product = models.Product{
		Code:              "DLT001",
		Name:              "Delta Banyo Bataryası",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Mix Bataryalar",
		ItemSubCategory:   "Delta Serisi",
		Origin:            "Çin",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/125/delta-banyo-bataryasi",
		PictureUrl:        "http://www.artiarmatur.com/images/product/aaf02f9b-a3d7-.jpg",
		UnitPrice:         290,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

	product = models.Product{
		Code:              "DLT002",
		Name:              "Delta Mix Uzun Musluk",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Mix Bataryalar",
		ItemSubCategory:   "Delta Serisi",
		Origin:            "Çin",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/132/delta-mix-uzun-musluk",
		PictureUrl:        "http://www.artiarmatur.com/images/product/b89684ea-b0f4-.jpg",
		UnitPrice:         90,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

	product = models.Product{
		Code:              "TES001",
		Name:              "Acem Uzun Musluk",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Tesisat",
		ItemSubCategory:   "Musluklar",
		Origin:            "Türkiye",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/1252/acem-uzun-musluk",
		PictureUrl:        "http://www.artiarmatur.com/images/product/6b08664a-a326-.jpg",
		UnitPrice:         170,
		CurrencyCode:      "TL",
	}

	product = models.Product{
		Code:              "TES002",
		Name:              "Osmanlı Musluğu",
		Brand:             "Artı",
		VAT:               18,
		ItemCategory:      "Tesisat",
		ItemSubCategory:   "Musluklar",
		Origin:            "Türkiye",
		BaseUnitOfMeasure: "ADET",
		Url:               "http://www.artiarmatur.com/details/1276/osmanli-muslugu-tugrali",
		PictureUrl:        "http://www.artiarmatur.com/images/product/b9487d42-37ce-.jpg",
		UnitPrice:         240,
		CurrencyCode:      "TL",
	}
	config.DB.Create(&product)
	fmt.Println(product.ID)

}
