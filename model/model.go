package model

import (
	"github.com/sankethkini/ConcurrencyInGo/config"
	"gorm.io/gorm"
)

var rates = config.LoadTaxConstants()

type Item interface {
	Calc() float64
}

type BaseItem struct {
	gorm.Model
	Name     string
	Price    float64
	Quantity int
	Total    float64
	Typ      string
}

func NewBaseItem(name string, price float64, typ string, quantity int) BaseItem {
	return BaseItem{Name: name, Price: price, Typ: typ, Quantity: quantity}
}

func (b *BaseItem) Calc() float64 {
	switch b.Typ {
	case "raw":
		b.Total = b.Price
		b.Total += b.Total * (rates.TaxRates.RawTax / 100)

	case "imported":
		b.Total = b.Price
		b.Total += b.Total * (rates.TaxRates.ImportTax / 100)
		switch {
		case b.Total <= 100:
			b.Total += rates.TaxRates.Surcharge100
		case b.Total <= 200:
			b.Total += rates.TaxRates.Surcharge200
		default:
			b.Total += b.Total * (rates.TaxRates.SurchargeMore / 100)
		}
	case "manufactured":
		b.Total = b.Price
		b.Total += b.Total * (rates.TaxRates.ManufacturedTax / 100)
		b.Total += b.Total * (rates.TaxRates.ManufacturedExtra / 100)

	default:
		b.Total = b.Price
	}
	return b.Total
}
