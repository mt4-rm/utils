package str

import (
	"fmt"
	"regexp"
	"strconv"
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

func ToInt64List(str string, sep string) ([]int64, error) {
	result := make([]int64, 0)
	strList := strings.Split(str, sep)
	for _, v := range strList {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, int64(intValue))
	}
	return result, nil
}

func ToIntList(str string, sep string) ([]int, error) {
	result := make([]int, 0)
	strList := strings.Split(str, sep)
	for _, v := range strList {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, intValue)
	}
	return result, nil
}

func IntListToStr(intList []int) string {
	result := ""
	for _, v := range intList {
		if len(result) > 0 {
			result += ","
		}
		result += strconv.Itoa(v)
	}
	return result
}

func Int64ListToStr(intList []int64) string {
	result := ""
	for _, v := range intList {
		if len(result) > 0 {
			result += ","
		}
		result += strconv.Itoa(int(v))
	}
	return result
}
