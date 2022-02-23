package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func SalesPricesApiRoot(api *gin.RouterGroup) {
	api.GET("/sales_prices", apicontrollers.GETSalesPrices)
	api.GET("/sales_prices/:id", apicontrollers.GETSalesPricesByIP)
	api.POST("/sales_prices", apicontrollers.POSTSalesPrices)
	api.PUT("/sales_prices/:id", apicontrollers.PUTSalesPrices)
	api.DELETE("/sales_prices/:id", apicontrollers.DELETESalesPrices)
}
