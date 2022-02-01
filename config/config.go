package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	DataBase struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

type TaxConstants struct {
	TaxRates struct {
		RawTax            float64 `yaml:"raw-tax"`
		ImportTax         float64 `yaml:"import-tax"`
		Surcharge100      float64 `yaml:"surcharge-100"`
		Surcharge200      float64 `yaml:"surcharge-200"`
		SurchargeMore     float64 `yaml:"surcharge-more"`
		ManufacturedTax   float64 `yaml:"manufactured-tax"`
		ManufacturedExtra float64 `yaml:"manufactured-extra"`
	} `yaml:"tax-rates"`
}

func LoadConfig() AppConfig {
	var config AppConfig
	err := cleanenv.ReadConfig("application.yaml", &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

func LoadTaxConstants() TaxConstants {
	var config TaxConstants

	err := cleanenv.ReadConfig("application.yaml", &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
