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

	TaxRates `yaml:"tax-rates"`
}

func LoadConfig() AppConfig {
	var config AppConfig
	err := cleanenv.ReadConfig("application.yaml", &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}
