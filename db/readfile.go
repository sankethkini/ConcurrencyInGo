package db

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sankethkini/ConcurrencyInGo/model"
)

func GetDataFromFile() ([]model.BaseItem, error) {
	fp, err := os.Open("data.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}
	var retData []model.BaseItem
	if len(data) == 0 {
		return retData, nil
	}
	err = json.Unmarshal(data, &retData)
	if err != nil {
		return nil, err
	}

	return retData, nil
}
