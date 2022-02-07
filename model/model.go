package model

import (
	"github.com/sankethkini/ConcurrencyInGo/config"
	"gorm.io/gorm"
)

type Item interface {
	Calc(config.TaxRates) float64
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

func (b *BaseItem) Calc(rates config.TaxRates) float64 {
	switch b.Typ {
	case "raw":
		b.Total = b.Price
		b.Total += b.Total * (rates.RawTax / 100)

	case "imported":
		b.Total = b.Price
		b.Total += b.Total * (rates.ImportTax / 100)
		switch {
		case b.Total <= 100:
			b.Total += rates.Surcharge100
		case b.Total <= 200:
			b.Total += rates.Surcharge200
		default:
			b.Total += b.Total * (rates.SurchargeMore / 100)
		}
	case "manufactured":
		b.Total = b.Price
		b.Total += b.Total * (rates.ManufacturedTax / 100)
		b.Total += b.Total * (rates.ManufacturedExtra / 100)

	default:
		b.Total = b.Price
	}
	return b.Total
}
