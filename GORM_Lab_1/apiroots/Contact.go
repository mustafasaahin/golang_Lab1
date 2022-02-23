package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func ContactApiRoot(api *gin.RouterGroup) {
	api.GET("/contact/", apicontrollers.GETContact)
	api.GET("/contact/:id", apicontrollers.GETContactByID)
	api.POST("/contact", apicontrollers.POSTContact)
	api.PUT("/contact/:id", apicontrollers.PUTContact)
	api.DELETE("/contact/:id", apicontrollers.DELETEContact)
}
