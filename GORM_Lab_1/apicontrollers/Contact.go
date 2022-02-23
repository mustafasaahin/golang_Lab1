package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

func GETContact(c *gin.Context) {
	var contact []models.Contact
	if err := config.DB.Find(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, contact)
}
func GETContactByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact models.Contact
	config.DB.
		Where("id=?", id).
		First(&contact)
	c.JSON(http.StatusOK, contact)
}
func POSTContact(c *gin.Context) {
	var contact models.Contact
	if err := c.Bind(&contact); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, contact)
		return
	}
}
func PUTContact(c *gin.Context) {
	var contact models.Contact
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&contact)
		config.DB.Save(&contact)
		c.JSON(http.StatusOK, contact)
		return
	}
}
func DELETEContact(c *gin.Context) {
	var contact models.Contact
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Delete(&contact).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, "Done")
		}
	}
}
