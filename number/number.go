package number

import (
	"math"
	"sort"

	timeUtils "github.com/mt4-rm/utils/time"
)

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

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func SortStrIntMapByValue(input map[string]int) []string {

	p := make(PairList, len(input))

	i := 0
	for k, v := range input {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted
	result := make([]string, len(p))
	for i, value := range p {
		result[i] = value.Key
	}
	return result
}

func SortDateStrFloat64Map(input map[string]float64) []string {
	result := make([]string, 0)
	for k := range input {
		result = append(result, k)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return timeUtils.DateStrToDate(result[i]).Before(timeUtils.DateStrToDate(result[j]))
	})
	return result
}
