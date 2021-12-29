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
		if col == "`GROUP`" || col == "BOOK" {
			for i, str := range parts {
				parts[i] = fmt.Sprintf("'%s'", str)
			}
		} else if col == "SYMBOL" {
			for i, str := range parts {
				parts[i] = fmt.Sprintf("%s like '%%%s%%'", col, str)
			}
			partsStr := strings.Join(parts, " OR ")
			return fmt.Sprintf("(%s)", partsStr)
		}
		partsStr := strings.Join(parts, ",")
		part = "(" + partsStr + ")"
	} else {
		if col == "SYMBOL" {
			return fmt.Sprintf("%s like '%%%s%%'", col, condition)
		} else {
			part = fmt.Sprintf("('%s')", condition)
		}

	}
	return fmt.Sprintf("%s in %s", col, part)
}
