package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func CustomerApiRoot(api *gin.RouterGroup) {
	api.GET("/customers", apicontrollers.GETCustomers)
	api.GET("/customer/:id", apicontrollers.GETCustomerByID)
	api.POST("/customers", apicontrollers.POSTCustomer)
	api.PUT("/customer/:id", apicontrollers.PUTCustomer)
	api.DELETE("/customer/:id", apicontrollers.DELETECustomer)
}
