package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestConverToMT4Date(t *testing.T) {
	location, _ := time.LoadLocation("Australia/Sydney")
	fmt.Println(time.Now())
	layout := "2006-01-02 15:04:05"
	tables := []struct {
		testTime time.Time
		answer   string
	}{{time.Date(
		2021, 8, 26, 23, 45, 26, 0, location), "2021-08-26 16:45:26"},
	}
	for _, unit := range tables {
		answer := ConverToMT4Date(unit.testTime).Format(layout)
		if answer != unit.answer {
			t.Errorf("util time test | convert to mt4 date | fail | %s | %s", unit.answer, answer)
		}
	}

}

func TestEpochToDate(t *testing.T) {
	tables := []struct {
		epoch  int64
		answer string
	}{
		{0, "1970-01-01 00:00:00"},
		{1632302747, "2021-09-22 09:25:47"},
	}
	for index, table := range tables {
		result := EpochToDate(table.epoch)
		if result.Format("2006-01-02 15:04:05") != table.answer {
			t.Errorf("test case %d failed. answer: %s expected: %s\n", index+1, result, table.answer)
		}
	}
}

func TestMTTime(t *testing.T) {
	fmt.Println(NewYorkTime())
}
