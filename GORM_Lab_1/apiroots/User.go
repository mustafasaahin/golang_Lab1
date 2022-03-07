package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/apicontrollers"
)

func UserApiRoot(api *gin.RouterGroup) {
	api.POST("/register", apicontrollers.POSTRegister)
	api.POST("/login", apicontrollers.POSTLogin)
}
