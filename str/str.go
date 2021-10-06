package str

import (
	"regexp"
	"strings"
)

func ExtractSymbolSuffix(symbol string) string {
	re := regexp.MustCompile(`[A-Z0-9]+`)
	main := re.FindString(symbol)
	suffix := strings.Replace(symbol, main, "", 1)
	return suffix
}
