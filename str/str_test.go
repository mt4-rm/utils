package str

import (
	"testing"
)

func TestTransformDotSeparateStrToSQLQuery(t *testing.T) {
	tables := []struct {
		col    string
		input  string
		answer string
	}{
		{
			"BOOK",
			"B,S,TEST",
			"BOOK in ('B','S','TEST')",
		},
		{
			"SYMBOL",
			"AUDUSD,XAUUSD,GBPJPY",
			"SYMBOL in ('AUDUSD','XAUUSD','GBPJPY')",
		},
		{
			"SYMBOL",
			"AUDUSD",
			"SYMBOL in ('AUDUSD')",
		},
		{
			"`GROUP`",
			"S_STD_USD,M_STD_USD",
			"`GROUP` in ('S_STD_USD','M_STD_USD')",
		},
		{
			"CMD",
			"1,0,2",
			"CMD in (1,0,2)",
		},
	}

	for _, table := range tables {
		result := TransformDotSeparateStrToSQLQuery(table.col, table.input)
		if result != table.answer {
			t.Error(result)
		}
	}
}
