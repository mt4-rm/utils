package number

import "math"

func RoundT4(number float64) float64 {
	return math.Round(number*10000) / 10000
}
func RoundT2(number float64) float64 {
	return math.Round(number*100) / 100
}

func FloorT2(number float64) float64 {
	return math.Floor(number*100) / 100
}

func CeilT2(number float64) float64 {
	return math.Ceil(number*100) / 100
}
