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

func TestSortDateStrMap(t *testing.T) {
	tables := []struct {
		input  map[string]float64
		answer []string
	}{
		{
			map[string]float64{
				"2022-01-03": 2.1,
				"2022-01-06": 2.1,
				"2022-01-10": 2.1,
				"2022-01-27": 2.1,
				"2022-01-19": 2.1,
				"2022-01-17": 2.1,
				"2022-02-01": 2.1,
				"2022-02-03": 2.1,
				"2022-02-04": 2.1,
				"1970-01-01": 2.1,
				"2022-02-28": 2.1,
				"2022-03-02": 2.1,
				"2022-03-03": 2.1,
				"2022-03-04": 2.1,
				"2022-03-07": 2.1,
				"2022-03-14": 2.1},
			[]string{
				"1970-01-01",
				"2022-01-03",
				"2022-01-06",
				"2022-01-10",
				"2022-01-17",
				"2022-01-19",
				"2022-01-27",
				"2022-02-01",
				"2022-02-03",
				"2022-02-04",
				"2022-02-28",
				"2022-03-02",
				"2022-03-03",
				"2022-03-04",
				"2022-03-07",
				"2022-03-14"},
		},
	}
	for _, table := range tables {
		result := SortDateStrFloat64Map(table.input)
		for i, v := range result {
			if table.answer[i] != v {
				t.Errorf("incorrect answer")
			}
		}
	}
}
