package main

import (
	"fmt"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/InitData"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/models"
)

func InitRecords() {
	if err := config.DB.Migrator().DropTable(
		models.Customer{},
		models.Contact{},
		models.Product{},
		models.SalesLine{},
		models.SalesHeader{},
		models.SalesPrice{},
	); err != nil {
		fmt.Println(err.Error())
	}

	if err := config.DB.AutoMigrate(
		models.Customer{},
		models.Contact{},
		models.Product{},
		models.SalesLine{},
		models.SalesHeader{},
		models.SalesPrice{}); err != nil {
		fmt.Println(err.Error())
	}
	InitData.InitCustomer()
	InitData.InitProduct()
	InitData.InitSalesPrice()
	InitData.InitSalesHeader()
}
