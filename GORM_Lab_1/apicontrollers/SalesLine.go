package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

func GETSalesLine(c *gin.Context) {
	var salesline []models.SalesLine
	if err := config.DB.
		Find(&salesline).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, salesline)
}
func GETSalesLineByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesline models.SalesLine
	config.DB.
		Where("id = ?", id).
		First(&salesline)
	c.JSON(http.StatusOK, salesline)
}
func POSTSalesLine(c *gin.Context) {
	var salesline models.SalesLine
	if err := c.Bind(&salesline); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&salesline).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, salesline)
		return
	}
}
func PUTSalesLine(c *gin.Context) {
	var salesline models.SalesLine
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&salesline).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		// Kaydı buldum
		_ = c.Bind(&salesline)
		config.DB.Save(&salesline)
		c.JSON(http.StatusOK, salesline)
		return
	}
}
func DELETESalesLine(c *gin.Context) {
	var salesline models.SalesLine
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&salesline).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Delete(&salesline).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, "Done")
		}
	}
}
