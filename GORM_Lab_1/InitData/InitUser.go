package InitData

import (
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"time"
)

func InitUser() {
	config.DB.Where("id<>0").
		Delete(&models.User{})
	var user = models.User{}
	user = models.User{
		Name:               "Mustafa",
		Surname:            "Sahin",
		Mail:               "shnmustafa28@gmail.com",
		Phone:              "05314246047",
		PersonalIdentityNo: "10895301298",
		BirthDate:          time.Date(1998, 2, 7, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Arti Armatur",
		Gender:             "Erkek",
		Password:           config.HashPassword("ms1998ms"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Kadir",
		Surname:            "Bilgi",
		Mail:               "kdrco@gmail.com",
		Phone:              "05367346285",
		PersonalIdentityNo: "10836272572",
		BirthDate:          time.Date(1999, 1, 16, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Kabil İnşaat",
		Gender:             "Erkek",
		Password:           config.HashPassword("kadir123"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Aleyna",
		Surname:            "Kaya",
		Mail:               "aleynakaya@gmail.com",
		Phone:              "05318402583",
		PersonalIdentityNo: "13668145945",
		BirthDate:          time.Date(1998, 7, 25, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Kabil İnşaat",
		Gender:             "Kadın",
		Password:           config.HashPassword("kaya98"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Meryem",
		Surname:            "Eren",
		Mail:               "erenmeryem24@gmail.com",
		Phone:              "05393386343",
		PersonalIdentityNo: "29032573450",
		BirthDate:          time.Date(1998, 12, 3, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Otin Mühendislik",
		Gender:             "Kadın",
		Password:           config.HashPassword("245738"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Mısra",
		Surname:            "Soylu",
		Mail:               "soylumisra@gmail.com",
		Phone:              "05417492729",
		PersonalIdentityNo: "28392515503",
		BirthDate:          time.Date(1998, 9, 16, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Akın İnşaat",
		Gender:             "Kadın",
		Password:           config.HashPassword("soylumisra98"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Sametcan",
		Surname:            "Yazici",
		Mail:               "lazsamo53@gmail.com",
		Phone:              "05457149855",
		PersonalIdentityNo: "16382639362",
		BirthDate:          time.Date(1998, 1, 16, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Şen Pres",
		Gender:             "Erkek",
		Password:           config.HashPassword("lazsamo53"),
	}
	config.DB.Create(&user)

	user = models.User{
		Name:               "Muhammet",
		Surname:            "Aksoy",
		Mail:               "aksoymuhammet28@gmail.com",
		Phone:              "05357301241",
		PersonalIdentityNo: "13475924128",
		BirthDate:          time.Date(1998, 11, 23, 0, 0, 0, 0, time.UTC),
		CompanyName:        "Seyisoglu",
		Gender:             "Erkek",
		Password:           config.HashPassword("563782"),
	}
	config.DB.Create(&user)
	user = models.User{
		Name:               "İbrahim",
		Surname:            "Cobani",
		Mail:               "ibrahim@imaconsult.com",
		Phone:              "05336292733",
		PersonalIdentityNo: "13475358128",
		BirthDate:          time.Date(1998, 11, 23, 0, 0, 0, 0, time.UTC),
		CompanyName:        "IciTech",
		Gender:             "Erkek",
		Password:           config.HashPassword("ibrahim_1975B1"),
	}
	config.DB.Create(&user)
}
