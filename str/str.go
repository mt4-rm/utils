package str

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractSymbolSuffix(symbol string) string {
	re := regexp.MustCompile(`[A-Z0-9]+`)
	main := re.FindString(symbol)
	suffix := strings.Replace(symbol, main, "", 1)
	return suffix
}

func TransformDotSeparateStrToSQLQuery(col, condition string) string {
	part := ""
	if strings.Contains(condition, ",") {
		parts := strings.Split(condition, ",")
		partsStr := strings.Join(parts, ",")
		part = "(" + partsStr + ")"
	} else {
		part = "(" + part + ")"
	}
	return fmt.Sprintf("%s in %s", col, part)
}
