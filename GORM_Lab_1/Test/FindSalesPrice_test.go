package Test

import (
	"fmt"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/SalesPriceManagement"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"testing"
)

func TestFinsSalesPrice(t *testing.T) {
	config.InitDB()
	customerNo := "MUS002"
	productCode := "PLS002"
	var customer models.Customer
	config.DB.Where("no = ?", customerNo).
		First(&customer)

	var product models.Product
	config.DB.Where("code = ?", productCode).
		First(&product)

	var salesPrices []models.SalesPrice
	config.DB.Where("customer_id = ? and product_id = ?", customer.ID, product.ID).
		Find(&salesPrices)

	unitPrice := SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		2,
		"ADET")

	fmt.Println(unitPrice)
}
