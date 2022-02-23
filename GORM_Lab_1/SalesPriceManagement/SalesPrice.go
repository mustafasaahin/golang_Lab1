package SalesPriceManagement

import "github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"

// FindSalesPrice fiyat listesinde ürüne ait uygun fiyatı bulan fonksiyon.
func FindSalesPrice(prices []models.SalesPrice, product models.Product, Quantity float64, UnitOfMeasure string) float64 {
	//Fiyat listesinde kayıt bulunmazsa liste fiıyatı üzerinden fiyatlandırılır
	if len(prices) == 0 {
		return product.UnitPrice
	}

	for _, price := range prices {
		if price.MinimumQuantity <= Quantity && price.UnitOfMeasure == UnitOfMeasure {
			return price.UnitPrice
		}
	}

	return product.UnitPrice

}
