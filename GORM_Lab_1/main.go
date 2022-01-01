package main

import (
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
)

func main() {
	config.InitDB()
	InitRecords()
}
