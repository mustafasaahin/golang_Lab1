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
		No:              "MUS001",
		Name:            "Seyisoğlu",
		ContactName:     "Ayhan Şahin",
		Phone:           "532 377 3480",
		Fax:             "0312 649 8344",
		Mail:            "info.ayhan.sahin@seyisoglu.com.tr",
		PaymentTerm:     1,
		CurrencyCode:    "Tl",
		Address:         "Yeşilpınar Mah Canan Sokak No 106 Eyüp/Istanbul",
		City:            "Ankara",
		Country:         "Turkey",
		TaxAreaCode:     "Çankaya",
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
		Phone:        "0312 556 6677",
		MobilPhone:   "0532 540 1194",
		Mail:         "ibrahim@seyisoglu.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Yasemin",
		LastName:     "KÖSE",
		Phone:        "0312 256 6677",
		MobilPhone:   "0532 520 1194",
		Mail:         "yasemin@seyisoglu.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ismail",
		LastName:     "ŞANSLI",
		Phone:        "0312 156 6677",
		MobilPhone:   "0532 521 1594",
		Mail:         "ismail@seyisoglu.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	customer = models.Customer{
		No:              "MUS002",
		Name:            "Artı Armatur",
		ContactName:     "Hakan Şahin",
		Phone:           "532 535 4909",
		Fax:             "0212 649 6327",
		Mail:            "hakan.sahin@seyisoglu.com.tr",
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

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Sevgi",
		LastName:     "TEMELSIZ",
		Phone:        "0211 526 6677",
		MobilPhone:   "0532 543 1194",
		Mail:         "sevgim@seyisoglu.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Yasemin",
		LastName:     "OTUR",
		Phone:        "0211 251 6677",
		MobilPhone:   "0532 525 1194",
		Mail:         "yasemin@seyisoglu.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Korkmaz",
		LastName:     "ŞANLI",
		Phone:        "0212 161 6617",
		MobilPhone:   "0532 501 1194",
		Mail:         "korkmaz@seyisoglu.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ismail",
		LastName:     "ALKAN",
		Phone:        "0212 611 6617",
		MobilPhone:   "0532 517 1994",
		Mail:         "ismail23@seyisoglu.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	customer = models.Customer{
		No:              "MUS003",
		Name:            "Akın İnşaat",
		ContactName:     "Süleyman CİGAL",
		Phone:           "532 327 2327",
		Fax:             "0212 425 6872",
		Mail:            "suleymancigal@akininsaat.com.tr",
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

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Mehmet",
		LastName:     "ALKAN",
		Phone:        "0210 566 6177",
		MobilPhone:   "0544 243 1794",
		Mail:         "mehmet@akininsaat.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ayşegül",
		LastName:     "ATILGAN",
		Phone:        "0211 111 6677",
		MobilPhone:   "0532 125 1794",
		Mail:         "aysegul@akininsaat.com.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ferdi",
		LastName:     "ILGIN",
		Phone:        "0212 151 6617",
		MobilPhone:   "0532 511 1174",
		Mail:         "ferdi@akininsaat.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Samet",
		LastName:     "GÜZEL",
		Phone:        "0211 651 6617",
		MobilPhone:   "0505 577 1964",
		Mail:         "sameet.guzel@akininsaat.com.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	customer = models.Customer{
		No:              "MUS004",
		Name:            "ŞenPres LTD ŞTİ",
		ContactName:     "Sedat Şahin",
		Phone:           "537 237 6238",
		Fax:             "0232 425 6872",
		Mail:            "sedatsahhin@senpres.com",
		PaymentTerm:     1,
		CurrencyCode:    "Tl",
		Address:         "Yeşilpınar Mah Begonya Sokak No 21 Eyüp/Istanbul",
		City:            "Izmir",
		Country:         "Turkey",
		TaxAreaCode:     "Alsancak",
		TaxNo:           "6347439",
		TradeRegisterNo: "235268",
	}
	config.DB.Create(&customer)
	fmt.Println(customer.ID)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Osman",
		LastName:     "ARICI",
		Phone:        "0232 161 6117",
		MobilPhone:   "0532 521 1374",
		Mail:         "osman.arici@senpres.com",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Yalçın",
		LastName:     "ATİK",
		Phone:        "0232 651 6617",
		MobilPhone:   "0505 577 1964",
		Mail:         "sameet.guzel@senpres.com",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ibrahim",
		LastName:     "ÇOBANİ",
		Phone:        "0232 556 6677",
		MobilPhone:   "0532 540 1194",
		Mail:         "ibrahim@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Sinan",
		LastName:     "TAYAK",
		Phone:        "0232 556 3452",
		MobilPhone:   "0532 345 3345",
		Mail:         "sinan@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Özlem",
		LastName:     "DEMİRCİOĞLU",
		Phone:        "0232 342 4568",
		MobilPhone:   "0532 122 3434",
		Mail:         "ozlem@senpres.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Burak",
		LastName:     "PALANDOKEN",
		Phone:        "0232 342 4312",
		MobilPhone:   "0532 122 6671",
		Mail:         "enes@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Samet",
		LastName:     "YAZICI",
		Phone:        "0232 971 2352",
		MobilPhone:   "0532 873 8732",
		Mail:         "ali@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Enes",
		LastName:     "YILDIRIM",
		Phone:        "0232 342 1291",
		MobilPhone:   "0532 122 6768",
		Mail:         "enes@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Ali",
		LastName:     "Çıtak",
		Phone:        "0232 971 7366",
		MobilPhone:   "0532 873 3494",
		Mail:         "ali@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Başak",
		LastName:     "Yılmaz",
		Phone:        "0232 813 7316",
		MobilPhone:   "0541 574 37 42",
		Mail:         "basak@senpres.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Meryem",
		LastName:     "EREN",
		Phone:        "0232 618 3261",
		MobilPhone:   "0532 359 6218",
		Mail:         "meryem@senpres.tr",
		Gender:       "Bayan",
	}
	config.DB.Create(&contact)

	contact = models.Contact{
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		FirstName:    "Mustafa",
		LastName:     "ŞAHİN",
		Phone:        "0232 649 8344",
		MobilPhone:   "05314246047",
		Mail:         "mustafa@senpres.tr",
		Gender:       "Bay",
	}
	config.DB.Create(&contact)

}
