package str

import (
	"testing"
)

func TestCheckRegex(t *testing.T) {
	tables := []struct {
		regex  string
		target string
		answer bool
	}{
		{
			`[A-Z0-9]+c$`,
			"XAUUSDc",
			true,
		},
		{
			`[A-Z0-9]+c$`,
			"XAUUSD-S",
			false,
		},
		{
			`[A-Z0-9]+c$`,
			"XAUUSD-c",
			false,
		},
		{
			`[A-Z0-9]+c$`,
			"US300c",
			true,
		},
	}
	for _, table := range tables {
		result := CheckRegex(table.regex, table.target)
		if table.answer != result {
			t.Errorf("check regex test fail - target: %s regex: %s result: %t answer: %t", table.target, table.regex, result, table.answer)
		}
	}
}
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
			`(SYMBOL like '%AUDUSD%' OR SYMBOL like '%XAUUSD%' OR SYMBOL like '%GBPJPY%')`,
		},
		{
			"SYMBOL",
			"AUDUSD",
			`SYMBOL like '%AUDUSD%'`,
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
