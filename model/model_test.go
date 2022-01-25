package model

import "testing"

func TestItems(t *testing.T) {
	var scrapData = []struct {
		Name     string
		Price    float64
		Quantity int
		Typ      string
		exp      float64
	}{
		{
			Name:     "abc",
			Price:    12300,
			Quantity: 12,
			Typ:      "raw",
			exp:      13837.5,
		},
		{
			Name:     "def",
			Price:    13000,
			Quantity: 12,
			Typ:      "imported",
			exp:      15015,
		},
		{
			Name:     "ghi",
			Price:    1400,
			Quantity: 33,
			Typ:      "manufactured",
			exp:      1606.5,
		},
		{
			Name:     "jkl",
			Price:    17800,
			Quantity: 22,
			Typ:      "raw",
			exp:      20025,
		},
		{
			Name:     "mno",
			Price:    13500,
			Quantity: 88,
			Typ:      "imported",
			exp:      15592.5,
		},
		{
			Name:     "pqr",
			Price:    80,
			Quantity: 33,
			Typ:      "imported",
			exp:      93,
		},
		{
			Name:     "stu",
			Price:    150,
			Quantity: 33,
			Typ:      "imported",
			exp:      175,
		},
	}

	for _, val := range scrapData {
		bs := NewBaseItem(val.Name, val.Price, val.Typ, val.Quantity)
		bs.Calc()
		_, _, _, total := bs.GetDetails()
		if total != val.exp {

			t.Errorf("inncorrect total calculation exp:%v got:%v", val.exp, total)
		}
	}
}
