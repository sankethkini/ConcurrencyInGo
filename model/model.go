package model

import "gorm.io/gorm"

const (
	RawTax float64 = 12.5

	ImportTax     float64 = 10
	Surcharge100  float64 = 5
	Surcharge200  float64 = 10
	SurchargeMore float64 = 5

	ManufacturedTax   float64 = 12.5
	ManufacturedExtra float64 = 2
)

//every type of item should implement this
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
		b.Total += b.Total * (RawTax / 100)

	case "imported":
		b.Total = b.Price
		b.Total += b.Total * (ImportTax / 100)
		if b.Total <= 100 {
			b.Total += Surcharge100
		} else if b.Total <= 200 {
			b.Total += Surcharge200
		} else {
			b.Total += b.Total * (SurchargeMore / 100)
		}
	case "manufactured":
		b.Total = b.Price
		b.Total += b.Total * (ManufacturedTax / 100)
		b.Total += b.Total * (ManufacturedExtra / 100)

	default:
		b.Total = b.Price
	}
	return b.Total
}
