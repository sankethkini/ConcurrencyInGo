package db

import (
	"fmt"

	"github.com/sankethkini/ConcurrencyInGo/config"
	"github.com/sankethkini/ConcurrencyInGo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//nolint: revive
//go:generate mockgen -destination db_mock.go -package db github.com/sankethkini/ConcurrencyInGo/db DBHelper
type DBHelper interface {
	ReadItems() ([]model.BaseItem, error)
}

type Client struct {
	db *gorm.DB
}

func (c *Client) IntializeDB() {
	conf := config.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DataBase.User, conf.DataBase.Password, conf.DataBase.Host, conf.DataBase.DBName)
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
		err = c.addData()
		if err != nil {
			panic(err)
		}
	}
}

func (c *Client) ReadItems() ([]model.BaseItem, error) {
	var items []model.BaseItem
	res := c.db.Find(&items)

	if res.Error != nil {
		return nil, res.Error
	}
	return items, nil
}

//nolint: gosec
func (c *Client) addData() error {
	data, err := GetDataFromFile()
	if err != nil {
		return err
	}
	for _, val := range data {
		c.db.Create(&val)
	}
	return nil
}

func NewClient() *Client {
	db := new(Client)
	db.IntializeDB()
	return db
}
