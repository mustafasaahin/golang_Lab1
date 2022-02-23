package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func SalesHeaderApiRoot(api *gin.RouterGroup) {
	api.GET("/sales_header", apicontrollers.GETSalesHeader)
	api.GET("/sales_header/:id", apicontrollers.GETSalesHeaderByID)
	api.POST("/sales_header", apicontrollers.POSTSalesHeader)
	api.PUT("/sales_header/:id", apicontrollers.PUTSalesHeader)
	api.DELETE("/sales_header/:id", apicontrollers.DELETESalesHeader)
}
