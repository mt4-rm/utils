package str

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ExtractMainAndEnding(symbol string) (string, string) {
	re := regexp.MustCompile(`[A-Z0-9]+`)
	main := re.FindString(symbol)
	ending := strings.Replace(symbol, main, "", 1)
	return main, ending
}

func ExtractMain(symbol string) string {
	re := regexp.MustCompile(`[A-Z0-9]+`)
	main := re.FindString(symbol)
	return main
}

func ConcatStr(quotesM map[string]bool) string {
	str := ""
	for k := range quotesM {
		str += k + ", "
	}
	str = strings.Trim(str, " ")

	str = strings.Trim(str, ",")
	return str
}

func GetCentMap() map[int]struct{} {
	centsStr := os.Getenv("CENT_SECURITIES")
	cents := strings.Split(centsStr, ",")
	centM := make(map[int]struct{})
	for _, cent := range cents {
		i, err := strconv.Atoi(cent)
		if err != nil {
			continue
		}
		centM[i] = struct{}{}
	}
	return centM
}
