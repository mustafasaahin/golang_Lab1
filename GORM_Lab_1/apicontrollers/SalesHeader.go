package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

func GETSalesHeader(c *gin.Context) {
	var salesheader []models.SalesHeader
	if err := config.DB.
		Find(&salesheader).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, salesheader)
}
func GETSalesHeaderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesheader models.SalesHeader
	config.DB.Where("id=?", id)
	c.JSON(http.StatusOK, salesheader)
	return
}
func POSTSalesHeader(c *gin.Context) {
	var salesheader models.SalesHeader
	if err := c.Bind(&salesheader); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&salesheader).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, salesheader)
		return
	}
}
func PUTSalesHeader(c *gin.Context) {
	var salesheader models.SalesHeader
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&salesheader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&salesheader)
		config.DB.Save(&salesheader)
		c.JSON(http.StatusOK, salesheader)
		return
	}
}
func DELETESalesHeader(c *gin.Context) {
	var salesheader models.SalesHeader
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&salesheader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Delete(&salesheader).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, "Done")
		}
	}
}
