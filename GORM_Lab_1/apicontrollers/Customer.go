package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/models"
	"net/http"
)

//GETCustomers: Müşteri listesini veren api controller
func GETCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := config.DB.
		Preload("Contacts").
		Find(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, customers)
}
