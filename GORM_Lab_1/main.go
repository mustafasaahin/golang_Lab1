package main

import (
	"fmt"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/InitData"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/models"
)

func main() {
	config.InitDB()
	if err := config.DB.AutoMigrate(
		&models.Book{},
		models.Car{},
		models.Customer{},
		models.Contact{},
		models.Product{},
		models.SalesLine{},
		models.SalesHeader{},
		models.SalesPrices{}); err != nil {
		fmt.Println(err.Error())
	}
	//InsertBookSample()
	//ListBooks()
	//ListCar()
	//InsertCarSample()
	//DeleteCar()
	InitData.InitCustomer()
	InitData.InitProduct()
}
func ListBooks() {
	var books []models.Book
	config.DB.Find(&books)
	for i, book := range books {
		fmt.Printf(
			"%d - ( %d ) Book Name : %s, Author: %s\n",
			i, book.ID, book.Name, book.Author)
	}
}
func ListCar() {
	var cars []models.Car
	config.DB.Find(&cars)
	for i, car := range cars {
		fmt.Printf(
			"%d - ( %d ) mark : %s, Nation :%s,Year : %s\n",
			i, car.ID, car.Mark, car.Nation, car.Year)
	}
}
func InsertBookSample() {
	var book models.Book
	book.Name = "Kar"
	book.Author = "Orhan PAMUK"
	config.DB.Create(&book)
}
func InsertCarSample() {
	var car models.Car
	car.Mark = "Renault"
	car.Year = "2014"
	car.Nation = "French"
	config.DB.Create(&car)

	car = models.Car{}
	car.Mark = "Toyota"
	car.Year = "2014"
	car.Nation = "Japan"
	config.DB.Create(&car)

	car = models.Car{}
	car.Mark = "BMW"
	car.Year = "2014"
	car.Nation = "Germany"
	config.DB.Create(&car)
}
func DeleteCar() {
	config.DB.
		//Where("year = ? and nation = ?",2014,"French").
		Delete(&models.Car{}, 2)
}
