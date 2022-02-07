//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/sankethkini/ConcurrencyInGo/application"
	"github.com/sankethkini/ConcurrencyInGo/config"
	"github.com/sankethkini/ConcurrencyInGo/db"
)

var appSet = wire.NewSet(
	wire.InterfaceValue(new(db.DBHelper), db.NewClient()),
	config.LoadConfig,
	config.NewRates,
	application.NewApp,
)

func IntializeApp() *application.MyApp {
	wire.Build(appSet)
	return &application.MyApp{}
}
