package config

//go:generate mockgen -destination tax_mock.go -package config github.com/sankethkini/ConcurrencyInGo/config ItaxRates
type ItaxRates interface {
	GetTaxRates() TaxRates
}

type TaxRates struct {
	RawTax            float64 `yaml:"raw-tax"`
	ImportTax         float64 `yaml:"import-tax"`
	Surcharge100      float64 `yaml:"surcharge-100"`
	Surcharge200      float64 `yaml:"surcharge-200"`
	SurchargeMore     float64 `yaml:"surcharge-more"`
	ManufacturedTax   float64 `yaml:"manufactured-tax"`
	ManufacturedExtra float64 `yaml:"manufactured-extra"`
}

type Rates struct {
	TaxRates
}

func NewRates(app AppConfig) ItaxRates {
	var rt Rates
	rt.TaxRates = app.TaxRates
	return &rt
}

func (r *Rates) GetTaxRates() TaxRates {
	return r.TaxRates
}
