package InitData

import (
	"fmt"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/SalesPriceManagement"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/config"
	"github.com/mustafasaahin/golang_Lab1/GORM_Lab_1/models"
	"math"
	"time"
)

func InitSalesHeader() {
	config.DB.
		Where("id <> 0").
		Delete(&models.SalesHeader{})
	config.DB.
		Where("id <> 0").
		Delete(&models.SalesLine{})

	var customer models.Customer
	config.DB.
		Where("no = ?", "MUS002").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader := models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		InvoiceNo:                  "12345",
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)

	var product models.Product
	config.DB.
		Where("code = ?", "PLS002").
		First(&product)
	var salesPrices []models.SalesPrice
	config.DB.Debug().Where("customer_id = ? and product_id = ?", customer.ID, product.ID).
		Find(&salesPrices)

	salesLine := models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitOfMeasure:  product.BaseUnitOfMeasure,
		VATRate:        product.VAT,
		Quantity:       5.0,
		DiscountAmount: 13.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	product = models.Product{}
	config.DB.
		Where("code = ?", "DLT001").
		First(&product)
	salesPrices = []models.SalesPrice{}
	config.DB.Debug().Where("customer_id = ? and product_id = ?", customer.ID, product.ID).
		Find(&salesPrices)

	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitOfMeasure:  product.BaseUnitOfMeasure,
		VATRate:        product.VAT,
		Quantity:       5.0,
		DiscountAmount: 13.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount
	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS003").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS003").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        product.VAT,
		Quantity:       7.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS004").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        1.0,
		Quantity:       7.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS002").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        18.0,
		Quantity:       15.0,
		DiscountAmount: 10.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS001").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        18.0,
		Quantity:       4.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS004").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        1.0,
		Quantity:       7.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS003").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        1.0,
		Quantity:       7.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

	customer = models.Customer{}
	config.DB.
		Where("no = ?", "MUS002").
		First(&customer)
	fmt.Println("Müşteri No :", customer.ID)

	salesHeader = models.SalesHeader{
		CustomerID:                 customer.ID,
		CustomerName:               customer.Name,
		CustomerContactName:        customer.ContactName,
		CustomerPhone:              customer.Phone,
		PostingDate:                time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		InvoiceNo:                  "12345",
		OrderDate:                  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		CustomerPersonalIdentityNo: customer.PersonalIdentityNo,
		CustomerFax:                customer.Fax,
		TaxAreCode:                 customer.TaxAreaCode,
		TaxNo:                      customer.TaxNo,
		CustomerMail:               customer.Mail,
		PaymentTerm:                customer.PaymentTerm,
		CustomerAddress:            customer.Address,
		CustomerCity:               customer.City,
		CustomerCountry:            customer.Country,
		DueDate:                    time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
	}
	config.DB.Create(&salesHeader)
	product = models.Product{}
	config.DB.
		Where("code = ?", "PLS001").
		First(&product)
	salesLine = models.SalesLine{
		SalesHeaderID:  salesHeader.ID,
		ProductID:      product.ID,
		ProductCode:    product.Code,
		ProductName:    product.Name,
		ProductBrand:   product.Brand,
		ProductOrigin:  product.Origin,
		UnitPrice:      product.UnitPrice,
		VATRate:        1.0,
		Quantity:       7.0,
		DiscountAmount: 15.0,
	}
	salesLine.UnitPrice = SalesPriceManagement.FindSalesPrice(
		salesPrices,
		product,
		salesLine.Quantity,
		"ADET")
	salesLine.LineAMount = (salesLine.Quantity * salesLine.UnitPrice) - salesLine.DiscountAmount
	salesLine.LineAMount = math.Floor(salesLine.LineAMount*100) / 100
	salesLine.VATAmount = math.Floor((salesLine.LineAMount*salesLine.VATRate/100)*100) / 100
	salesLine.Amount = salesLine.LineAMount + salesLine.VATAmount
	salesLine.Amount = math.Floor(salesLine.Amount*100) / 100
	config.DB.Create(&salesLine)

	salesHeader.TotalVAT += salesLine.VATAmount
	salesHeader.TotalAmount += salesLine.Amount

	config.DB.Save(&salesHeader)

}
