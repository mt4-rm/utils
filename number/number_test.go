package number

import (
	"testing"
)

func TestSortStrIntMap(t *testing.T) {
	tables := []struct {
		input  map[string]int
		answer []string
	}{
		{
			map[string]int{"1": 123, "2": 1, "3": -12, "4": 50},
			[]string{"1", "4", "2", "3"},
		},
	}
	for _, table := range tables {
		result := SortStrIntMapByValue(table.input)
		for i, v := range result {
			if table.answer[i] != v {
				t.Errorf("incorrect answer")
			}
		}
	}
}
