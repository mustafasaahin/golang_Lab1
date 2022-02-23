package InitData

import (
	"fmt"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"math"
	"time"
)

func InitSalesPrice() {
	fmt.Println("")
	fmt.Println("Fiyat Listesi Ekleniyor..")
	fmt.Println("-------------------------")
	var customer models.Customer
	discountPercentage := 1.1
	config.DB.
		Where("no = ?", "MUS001").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	var product models.Product
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)

	salesPrice := models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS003").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "TES002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS002").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)
	discountPercentage = 1.15
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS003").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)
	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "TES002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS003").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)
	discountPercentage = 1.08
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS003").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "TES002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS004").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)
	discountPercentage = 1.2
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS003").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT001").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

	product = models.Product{}
	config.DB.
		Where("code = ?", "TES002").
		First(&product)

	salesPrice = models.SalesPrice{
		CustomerID:      customer.ID,
		CustomerNo:      customer.No,
		CustomerName:    customer.Name,
		ProductID:       product.ID,
		ProductCode:     product.Code,
		ProductName:     product.Name,
		UnitPrice:       math.Floor(product.UnitPrice/discountPercentage*100) / 100,
		CurrencyCode:    "TL",
		MinimumQuantity: 5,
		UnitOfMeasure:   "ADET",
		StartDate:       time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesPrice)

}
