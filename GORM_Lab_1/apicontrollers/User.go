package apicontrollers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/libs"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"net/http"
)

func POSTRegister(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	form.Password = config.HashPassword(form.Password)
	if err := config.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
		return
	}

}

func POSTLogin(c *gin.Context) {
	type LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var form LoginForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(form.Username, form.Password)
	var user models.User
	if err := config.DB.Debug().Where("mail=? and password=?", form.Username,
		config.HashPassword(form.Password)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	if token, err := libs.CreateToken(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		exp := libs.ParseToken(token, "expd")

		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"token":   token,
			"exp":     exp,
			"name":    user.Name,
			"surname": user.Surname,
			"phone":   user.Phone,
		})
		return
	}
}
