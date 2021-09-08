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

func TestMTTime(t *testing.T) {
	fmt.Println(NewYorkTime())
}
