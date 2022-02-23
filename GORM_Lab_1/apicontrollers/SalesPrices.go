package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

func GETSalesPrices(c *gin.Context) {
	var salesprices []models.SalesPrice
	if err := config.DB.
		Find(&salesprices).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, salesprices)
}
func GETSalesPricesByIP(c *gin.Context) {
	id := c.Params.ByName("id")
	var salesprices models.SalesPrice
	config.DB.Where("id", id).
		First(&salesprices)
	c.JSON(http.StatusOK, salesprices)
}
func POSTSalesPrices(c *gin.Context) {
	var salesprices models.SalesPrice
	if err := c.Bind(&salesprices); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&salesprices).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, salesprices)
		return
	}
}
func PUTSalesPrices(c *gin.Context) {
	var salesprices models.SalesPrice
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&salesprices).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&salesprices)
		config.DB.Save(&salesprices)
		c.JSON(http.StatusOK, salesprices)
		return
	}
}
func DELETESalesPrices(c *gin.Context) {
	var salesprices models.SalesPrice
	var product models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&salesprices).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		config.DB.Delete(&product)
		if err := config.DB.Delete(&salesprices).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, "Done")
		}
	}
}
