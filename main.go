package main

import (
	"github.com/sankethkini/ConcurrencyInGo/application"
	"github.com/sankethkini/ConcurrencyInGo/db"
)

func main() {
	ap := application.NewApp(db.NewClient())
	ap.Start()
}
