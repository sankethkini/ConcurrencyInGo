package main

import (
	"github.com/sankethkini/ConcurrencyInGo/provider"
)

func main() {
	ap := provider.IntializeApp()
	ap.Start()
}
