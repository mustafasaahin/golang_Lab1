package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

func GETProducts(c *gin.Context) {
	var products []models.Product
	filter := make(map[string]interface{})

	if itemCategoryFilter, itemCategoryFilterExist := c.GetQuery("item_category"); itemCategoryFilterExist {
		filter["item_category"] = itemCategoryFilter
	}

	if originFilter, originFilterExist := c.GetQuery("origin"); originFilterExist {
		filter["origin"] = originFilter
	}
	config.DB.
		Preload("SalesPrices").
		Where(filter).
		Find(&products)

	c.JSON(http.StatusOK, products)
}

func GETProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	config.DB.
		Preload("SalesPrices").
		Where("id = ?", id).
		First(&product)
	c.JSON(http.StatusOK, product)
}

func POSTProduct(c *gin.Context) {
	var form models.Product
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, form)
		return
	}
}

func PUTProduct(c *gin.Context) {
	var form models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		// Kaydı buldum
		_ = c.Bind(&form)
		config.DB.Save(&form)
		c.JSON(http.StatusOK, form)
		return
	}
}

func DELETEProduct(c *gin.Context) {
	var form models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		// Kaydı bulduk.
		//TODO:Eğer ürün Sales line yada sales price tablosunda kullanılmış ise silmesine izin vermemek lazım
		if err := config.DB.Delete(&form).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, "Done")
		}
	}
}
