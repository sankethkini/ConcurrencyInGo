package application

import (
	"fmt"
	"sync"

	"github.com/sankethkini/ConcurrencyInGo/db"
	"github.com/sankethkini/ConcurrencyInGo/model"
)

type MyApp struct {
	sqlClient  db.IClinet
	items      []model.BaseItem
	finalItems []model.BaseItem
	wt         sync.WaitGroup
	itemMutex  sync.Mutex
	finalMutex sync.Mutex
}

func NewApp(db db.IClinet) *MyApp {
	app := MyApp{sqlClient: db}
	return &app
}

func (app *MyApp) Start() {
	rows, err := app.sqlClient.ReadItems()
	if err != nil {
		fmt.Println("cannot read db %w", err)
		return
	}
	app.wt.Add(1)
	out := app.AddToList(rows)
	out1 := app.CalcTotal(out)
	out2 := app.ItemToCollection(out1)
	app.DisplayItems(out2)
	app.wt.Wait()
}

func (app *MyApp) AddToList(items []model.BaseItem) chan model.BaseItem {
	out := make(chan model.BaseItem)
	go func() {
		defer close(out)
		for _, it := range items {
			app.itemMutex.Lock()
			app.items = append(app.items, it)
			app.itemMutex.Unlock()
			out <- it
		}
	}()
	return out
}

func (app *MyApp) CalcTotal(in chan model.BaseItem) chan model.BaseItem {
	out := make(chan model.BaseItem)
	go func() {
		defer close(out)
		for val := range in {
			val.Calc()
			out <- val
		}
	}()

	return out
}

func (app *MyApp) ItemToCollection(ch chan model.BaseItem) chan model.BaseItem {
	out := make(chan model.BaseItem)
	go func() {
		defer close(out)
		for mt := range ch {
			app.finalMutex.Lock()
			app.finalItems = append(app.finalItems, mt)
			app.finalMutex.Unlock()
			out <- mt
		}
	}()
	return out
}

func (app *MyApp) DisplayItems(ch <-chan model.BaseItem) {
	go func() {
		for val := range ch {
			fmt.Printf("Name: %s Price: %v quantity: %d  total: %v \n", val.Name, val.Price, val.Quantity, val.Total)
		}
		app.wt.Done()
	}()
}
