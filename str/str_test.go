package str

import (
	"fmt"
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
			"'B','S','TEST'",
			"BOOK in ('B','S','TEST')",
		},
		{
			"CMD",
			"'1','0','2'",
			"CMD in ('1','0','2')",
		},
	}

	for _, table := range tables {
		result := TransformDotSeparateStrToSQLQuery(table.col, table.input)
		if result != table.answer {
			fmt.Println(result)
		}
	}
}
