package model

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sankethkini/ConcurrencyInGo/config"
)

func TestItems(t *testing.T) {
	scrapData := []struct {
		Name     string
		Price    float64
		Quantity int
		Typ      string
		exp      float64
	}{
		{
			Name:     "abc",
			Price:    12300,
			Quantity: 12,
			Typ:      "raw",
			exp:      13837.5,
		},
		{
			Name:     "def",
			Price:    13000,
			Quantity: 12,
			Typ:      "imported",
			exp:      15015,
		},
		{
			Name:     "ghi",
			Price:    1400,
			Quantity: 33,
			Typ:      "manufactured",
			exp:      1606.5,
		},
		{
			Name:     "jkl",
			Price:    17800,
			Quantity: 22,
			Typ:      "raw",
			exp:      20025,
		},
		{
			Name:     "mno",
			Price:    13500,
			Quantity: 88,
			Typ:      "imported",
			exp:      15592.5,
		},
		{
			Name:     "pqr",
			Price:    80,
			Quantity: 33,
			Typ:      "imported",
			exp:      93,
		},
		{
			Name:     "stu",
			Price:    150,
			Quantity: 33,
			Typ:      "imported",
			exp:      175,
		},
		{
			Name:     "some",
			Price:    150,
			Quantity: 33,
			Typ:      "some",
			exp:      150,
		},
	}

	tax := config.TaxRates{
		RawTax:            12.5,
		ImportTax:         10,
		ManufacturedExtra: 2,
		Surcharge100:      5,
		Surcharge200:      10,
		SurchargeMore:     5,
		ManufacturedTax:   12.5,
	}

	for _, val := range scrapData {
		bs := NewBaseItem(val.Name, val.Price, val.Typ, val.Quantity)
		ctrl := gomock.NewController(t)
		cfg := config.NewMockItaxRates(ctrl)
		cfg.EXPECT().GetTaxRates().Return(tax)
		bs.Calc(cfg.GetTaxRates())
		if bs.Total != val.exp {
			t.Errorf("inncorrect total calculation exp:%v got:%v", val.exp, bs.Total)
		}
	}
}
