package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"gorm.io/gorm"
	"net/http"
)

//GETCustomers: Müşteri listesini veren api controller
func GETCustomers(c *gin.Context) {
	var customers []models.Customer
	if err := config.DB.
		Preload("Contacts").
		Find(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, customers)
		return
	}
}
func GETCustomerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var customer models.Customer
	if err := config.DB.Where("id =?", id).First(&customer).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, customer)
}

func POSTCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.Bind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, customer)
		return
	}
}
func PUTCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Params.ByName("id")
	config.DB.
		Preload("Customer").
		Where("id=?", id).
		First(&customer)
	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		// Kaydı buldum
		_ = c.Bind(&customer)
		config.DB.Save(&customer)
		c.JSON(http.StatusOK, customer)
		return
	}
}
func DELETECustomer(c *gin.Context) {
	var customer models.Customer
	var contact models.Contact
	var salesprices models.SalesPrice
	var salesheader models.SalesHeader
	id := c.Params.ByName("id")
	config.DB.
		Preload("Customer").
		Where("id=?", id).
		First(&customer)

	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		//customer, customerID üzerinden salesprice , salesheader ve contact'ta kullanılıyor.Bu yüzden;
		//1-salesprice sil dedik. Hata çıkarsa hata kodu döndür.
		//2-salesheader sil dedik. Hata çıkarsa hata kodu döndür.
		//3-contact sil dedik. Hata çıkarsa hata kodu döndür.
		//4-salesprice ,salesheader ve contact başarılı şekilde silinmişse customer sil.Hata çıkarsa hata kodu döndür.
		//5-Hata çıkmazsa "Done" yazdır ve sil.
		if err := config.DB.Where("customer_id=?", customer.ID).Delete(&salesprices).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := config.DB.Where("customer_id=?", customer.ID).Delete(&salesheader).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := config.DB.Where("customer_id", customer.ID).Delete(&contact).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := config.DB.Where("id=?", customer.ID).Delete(&customer).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return

		} else {
			c.JSON(http.StatusOK, "Done")
		}

		/*	if err := config.DB.Where("customer_id = ?", customer.ID).Delete(&contact).Error; err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			if err := config.DB.Delete(&customer).Error; err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			} else {
				c.JSON(http.StatusOK, "Done")
			}
		} */
	}
}
