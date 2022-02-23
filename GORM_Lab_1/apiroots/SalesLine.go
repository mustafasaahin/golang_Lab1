package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func SalesLineApiRoot(api *gin.RouterGroup) {
	api.GET("/sales_line", apicontrollers.GETSalesLine)
	api.GET("/sales_line/:id", apicontrollers.GETSalesLineByID)
	api.POST("/sales_line", apicontrollers.POSTSalesLine)
	api.PUT("/sales_line/:id", apicontrollers.PUTSalesLine)
	api.DELETE("/sales_line/:id", apicontrollers.DELETESalesLine)
}
