package application

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sankethkini/ConcurrencyInGo/config"
	"github.com/sankethkini/ConcurrencyInGo/db"
	"github.com/sankethkini/ConcurrencyInGo/model"
)

var ScrapData = []model.BaseItem{
	{
		Name:     "abc",
		Price:    12300,
		Quantity: 12,
		Typ:      "raw",
	},
	{
		Name:     "def",
		Price:    13000,
		Quantity: 12,
		Typ:      "imported",
	},
	{
		Name:     "ghi",
		Price:    1400,
		Quantity: 33,
		Typ:      "manufactured",
	},
	{
		Name:     "jkl",
		Price:    17800,
		Quantity: 22,
		Typ:      "raw",
	},
	{
		Name:     "mno",
		Price:    13500,
		Quantity: 88,
		Typ:      "imported",
	},
}

func TestDisplay(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := db.NewMockDBHelper(ctrl)
	mockRates := config.NewMockItaxRates(ctrl)
	mock.EXPECT().ReadItems().Times(1).Return(ScrapData, nil)
	mockRates.EXPECT().GetTaxRates().Times(5).Return(config.TaxRates{RawTax: 12.5, ImportTax: 10, Surcharge100: 5, Surcharge200: 10, SurchargeMore: 5, ManufacturedTax: 12.5, ManufacturedExtra: 2})

	app := NewApp(mock, mockRates)
	app.Start()
	for i := range app.finalItems {
		if app.finalItems[i].Name != ScrapData[i].Name {
			t.Errorf("not equal")
		}
	}
	for i := range app.items {
		if app.items[i].Name != ScrapData[i].Name {
			t.Errorf("not equal")
		}
	}
}
