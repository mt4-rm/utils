package time

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"4d63.com/tz"
)

func ConverToMT4Date(now time.Time) time.Time {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	return now.In(location).Add(time.Hour * 7)
}

func MT4DateEpoch() int64 {
	mt4Now := MTTime()
	today := time.Date(mt4Now.Year(), mt4Now.Month(), mt4Now.Day(), 0, 0, 0, 0, time.UTC)
	return today.Unix()
}

func GetLocaltion() *time.Location {
	locationStr := os.Getenv("LOCATION")
	location, err := time.LoadLocation(locationStr)
	if err != nil {
		log.Fatal().Err(err).Msgf("load %s location fail", locationStr)
	}
	return location
}

func NewYorkTime() time.Time {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	t := time.Now()
	return t.In(location)
}

func MTTime() time.Time {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	t := time.Now()
	return t.In(location).Add(time.Hour * 7)
}

func MTTimeStr() string {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	t := time.Now()
	return t.In(location).Add(time.Hour * 7).Format("2006-01-02 15:04:05")
}

func MTEpoch() int64 {
	_, off := MTTime().Zone()
	return MTTime().Unix() + int64(off)
}

func EpochToTime(epoch int64) time.Time {
	return time.Unix(epoch, 0).UTC()
}

func EpochToTimeStr(epoch int64) string {
	return time.Unix(epoch, 0).UTC().Format("2006-01-02 15:04:05")
}

func EpochToDateStr(epoch int64) string {
	return time.Unix(epoch, 0).UTC().Format("2006-01-02")
}

// epoch of today date
func MTDateEpoch() int64 {
	layout := "2006-01-02"
	currentDate := MTDate().Format(layout)
	t, err := time.Parse(layout, currentDate)
	if err != nil {
		log.Fatal().Err(err).Msgf("utils | parse date string to date | fail | %s", currentDate)
	}
	return t.Unix()
}

func MTDate() time.Time {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	t := time.Now()
	t = t.In(location).Add(time.Hour * 7)
	returnDate, _ := time.Parse("2006-01-02", t.Format("2006-01-02"))
	return returnDate
}

func MTDateStr() string {
	location, errt := tz.LoadLocation("America/New_York")
	if errt != nil {
		fmt.Println(errt)
	}
	t := time.Now()
	return t.In(location).Add(time.Hour * 7).Format("2006-01-02")
}

func RemoveTimeZone(t time.Time) time.Time {
	str := t.Format("2006-01-02 15:04:05")
	new, _ := time.Parse("2006-01-02 15:04:05", str)
	return new
}

func GetRelativeEpoch(hour, min, sec int32) int32 {
	return hour*3600 + min*60 + sec
}

func TrimEpochDate(epoch int64) int {
	return int(epoch) % 86400
}

func StripEpochTime(epoch int64) int64 {
	return epoch - epoch%86400
}

func SameDay(date time.Time, year, month, day int32) bool {
	if date.Year() == int(year) && int(date.Month()) == int(month) && date.Day() == int(day) {
		return true
	}
	return false
}

func DateStrToDate(date string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}
	return t
}

func TimeStrToTime(date string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}
	return t
}

func TimeStrToEpoch(date string) int64 {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
	}
	return t.Unix()
}

// func TimeToTimeStr(date time.Time) string {

// }

func TimeToEpoch(date time.Time) int64 {
	str := date.Format("2006-01-02 15:04:05")
	return TimeStrToEpoch(str)
}

func InTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

func HMSStrToHMS(timeStr string) time.Time {
	newLayout := "15:04"
	result, err := time.Parse(newLayout, timeStr)
	if err != nil {
		fmt.Println("can not parse  ", timeStr)
	}
	return result
}
