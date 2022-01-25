package application

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sankethkini/ConcurrencyInGo/db"
)

//TODO
func TestDisplay(t *testing.T) {

	ctrl := gomock.NewController(t)
	mock := db.NewMockDBHelper(ctrl)
	mock.EXPECT().ReadDB().Times(1).Return(db.ScrapData, nil)

	app := NewApp(mock)
	app.Start()
	//TODO
}
