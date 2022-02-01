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

	var rat config.AppConfig
	rat.TaxRates.RawTax = 12.5
	rat.TaxRates.ImportTax = 10
	rat.TaxRates.ManufacturedExtra = 2
	rat.TaxRates.Surcharge100 = 5
	rat.TaxRates.Surcharge200 = 10
	rat.TaxRates.SurchargeMore = 5
	rat.TaxRates.ManufacturedTax = 12.5

	for _, val := range scrapData {
		bs := NewBaseItem(val.Name, val.Price, val.Typ, val.Quantity)
		ctrl := gomock.NewController(t)
		cfg := config.NewMockIConfig(ctrl)
		cfg.EXPECT().LoadConfig().Return(rat)
		bs.Calc(cfg.LoadConfig())
		if bs.Total != val.exp {
			t.Errorf("inncorrect total calculation exp:%v got:%v", val.exp, bs.Total)
		}
	}
}
