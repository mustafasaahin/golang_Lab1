package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
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
	var product models.Product
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, product)
		return
	}
}
func PUTProduct(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		// Kaydı buldum
		_ = c.Bind(&product)
		config.DB.Save(&product)
		c.JSON(http.StatusOK, product)
		return
	}
}
func DELETEProduct(c *gin.Context) {
	var product models.Product
	var salesprices models.SalesPrice
	var salesline models.SalesLine
	id := c.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		// Kaydı bulduk.
		//product, productID üzerinden salesprice ve salesline'da kullanılıyor.Bu yüzden;
		//1-salesprice sil dedik. Hata çıkarsa hata kodu döndür.
		//2-salesline sil dedik. Hata çıkarsa hata kodu döndür.
		//3-salesprice ve salesline başarılı şekilde silinmişse prodcut sil.Hata çıkarsa hata kodu döndür.
		//4-Hata çıkmazsa "Done" yazdır ve sil.

		if err := config.DB.Where("product_id=?", product.ID).Delete(&salesprices).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := config.DB.Where("product_id=?", product.ID).Delete(&salesline).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := config.DB.Where("id=?", product.ID).Delete(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Done")

		//TODO:Eğer ürün Sales line yada sales price tablosunda kullanılmış ise silmesine izin vermemek lazım
		/*	if err := config.DB.Delete(&product).Error; err != nil {
				if err := config.DB.Where("id =?", id).Delete(&salesprices).Error; err != nil {
					c.JSON(http.StatusBadRequest, err.Error())
					return
				}
			} else {
				c.JSON(http.StatusOK, "Done")
			} */

	}
}
