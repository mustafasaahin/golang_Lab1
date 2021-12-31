package InitData

import (
	"fmt"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/models"
)

func InitCustomer() {
	config.DB.
		Where("id <> 0").
		Delete(&models.Customer{})

	config.DB.
		Where("id <> 0").
		Delete(&models.Contact{})

	var customer = models.Customer{

		Name:            "Seyisoğlu",
		ContactName:     "Ayhan Şahin",
		Phone:           "532 377 3480",
		Fax:             "0212 649 8344",
		Mail:            "info.ayhan.sahin@gmail.com",
		PaymentTerm:     1,
		CurrencyCode:    "Tl",
		Address:         "Yeşilpınar Mah Canan Sokak No 106 Eyüp/Istanbul",
		City:            "Istanbul",
		Country:         "Turkey",
		TaxAreaCode:     "Eyüp",
		TaxNo:           "634789",
		TradeRegisterNo: "552637",
	}
	config.DB.Create(&customer)
	fmt.Println(customer.ID)

	contact := models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Osman",
		LastName:     "Canli",
		Phone:        "0212 556 6677",
		MobilPhone:   "0532 540 1194",
		Mail:         "ibrahim@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	customer = models.Customer{
		Name:            "Artı Armatur",
		ContactName:     "Hakan Şahin",
		Phone:           "532 535 4909",
		Fax:             "0212 649 6327",
		Mail:            "hakan.sahin@gmail.com",
		PaymentTerm:     1,
		CurrencyCode:    "Tl/Usd",
		Address:         "Demirören Mah İklim Sokak No 45 Avcılar/Istanbul",
		City:            "Istanbul",
		Country:         "Turkey",
		TaxAreaCode:     "Avcılar",
		TaxNo:           "642672",
		TradeRegisterNo: "98273",
	}

	config.DB.Create(&customer)
	fmt.Println(customer.ID)
	customer = models.Customer{
		Name:            "Akın İnşaat",
		ContactName:     "Süleyman CİGAL",
		Phone:           "532 327 2327",
		Fax:             "0212 425 6872",
		Mail:            "suleymancigal@gmail.com",
		PaymentTerm:     1,
		CurrencyCode:    "Tl",
		Address:         "Paşa Mah Gül Sokak No 21 Beylikdüzü/Istanbul",
		City:            "Istanbul",
		Country:         "Turkey",
		TaxAreaCode:     "Beylikdüzü",
		TaxNo:           "6347439",
		TradeRegisterNo: "235268",
	}
	config.DB.Create(&customer)
	fmt.Println(customer.ID)

	customer = models.Customer{
		Name:            "ŞenPres LTD ŞTİ",
		ContactName:     "Sedat Şahin",
		Phone:           "537 237 6238",
		Fax:             "0212 425 6872",
		Mail:            "sedatsahhin@gmail.com",
		PaymentTerm:     1,
		CurrencyCode:    "Tl",
		Address:         "Yeşilpınar Mah Begonya Sokak No 21 Eyüp/Istanbul",
		City:            "Istanbul",
		Country:         "Turkey",
		TaxAreaCode:     "Eyüp",
		TaxNo:           "6347439",
		TradeRegisterNo: "235268",
	}
	config.DB.Create(&customer)
	fmt.Println(customer.ID)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ibrahim",
		LastName:     "ÇOBANİ",
		Phone:        "0212 556 6677",
		MobilPhone:   "0532 540 1194",
		Mail:         "ibrahim@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Sinan",
		LastName:     "TAYAK",
		Phone:        "0212 556 3452",
		MobilPhone:   "0532 345 3345",
		Mail:         "sinan@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Özlem",
		LastName:     "DEMİRCİOĞLU",
		Phone:        "0212 342 4568",
		MobilPhone:   "0532 122 3434",
		Mail:         "ozlem@b1.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Burak",
		LastName:     "PALANDOKEN",
		Phone:        "0212 342 4312",
		MobilPhone:   "0532 122 6671",
		Mail:         "enes@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Samet",
		LastName:     "YAZICI",
		Phone:        "0212 971 2352",
		MobilPhone:   "0532 873 8732",
		Mail:         "ali@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Enes",
		LastName:     "YILDIRIM",
		Phone:        "0212 342 1291",
		MobilPhone:   "0532 122 6768",
		Mail:         "enes@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ali",
		LastName:     "Çıtak",
		Phone:        "0212 971 7366",
		MobilPhone:   "0532 873 3494",
		Mail:         "ali@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Başak",
		LastName:     "Yılmaz",
		Phone:        "0212 813 7316",
		MobilPhone:   "0541 574 37 42",
		Mail:         "basak@b1.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Meryem",
		LastName:     "EREN",
		Phone:        "0212 618 3261",
		MobilPhone:   "0532 359 6218",
		Mail:         "meryem@b1.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)
	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Mustafa",
		LastName:     "ŞAHİN",
		Phone:        "0212 649 8344",
		MobilPhone:   "05314246047",
		Mail:         "mustafa@b1.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

}
