package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apiroots"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"net/http"
)

func main() {
	config.InitDB()
	InitRecords()

	app := gin.Default()
	api := app.Group("/api")
	apiroots.ContactApiRoot(api)
	apiroots.ProductApiRoot(api)
	apiroots.SalesHeaderApiRoot(api)
	apiroots.SalesLineApiRoot(api)
	apiroots.SalesPricesApiRoot(api)
	apiroots.CustomerApiRoot(api)
	apiroots.UserApiRoot(api)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Minnak Kurbaa",
			"duduk":   "findı",
		})
	})
	app.Run(config.GetOutboundIP() + ":8014")
}
