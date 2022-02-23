package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func ProductApiRoot(api *gin.RouterGroup) {
	api.GET("/products", apicontrollers.GETProducts)
	api.GET("/product/:id", apicontrollers.GETProductByID)
	api.POST("/product", apicontrollers.POSTProduct)
	api.PUT("/product/:id", apicontrollers.PUTProduct)
	api.DELETE("/product/:id", apicontrollers.DELETEProduct)
}
