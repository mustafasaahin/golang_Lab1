package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/apicontrollers"
)

func CustomerApiRoot(api *gin.RouterGroup) {
	api.GET("/customers", apicontrollers.GETCustomers)
}
