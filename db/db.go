package db

import (
	"github.com/sankethkini/ConcurrencyInGo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TableName = "items"
	DbName    = "itemdb"
	Password  = "1234"
	UserName  = "root"
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

type DBHelper interface {
	ReadDB() ([]model.BaseItem, error)
}

type Client struct {
	db *gorm.DB
}

func (c *Client) IntializeDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/itemdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	c.db = db
	if !db.Migrator().HasTable(&model.BaseItem{}) {
		err := c.db.AutoMigrate(&model.BaseItem{})
		if err != nil {
			panic(err)
		}
		c.addData()
	}

}

func (c *Client) ReadDB() ([]model.BaseItem, error) {
	var items []model.BaseItem
	res := c.db.Find(&items)
	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

func (c *Client) addData() error {

	for _, val := range ScrapData {
		c.db.Create(&val)
	}
	return nil
}

func NewClient() *Client {
	db := Client{}
	db.IntializeDB()
	return &db
}
