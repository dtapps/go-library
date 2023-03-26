package rocron

import "fmt"

func concatenateStrings(day int) string {
	var result string
	if day > 1 {
		result += "1"
	}
	for i := 1; i <= 31; i++ {
		if isDivisibleBy(i, day) {
			if len(result) > 0 {
				result += ","
			}
			result += fmt.Sprintf("%d", i)
		}
	}

	return result
}

func isDivisibleBy(num, divisor int) bool {
	return num%divisor == 0
}

func removeFirstTwoChars(str string) string {
	if len(str) > 2 {
		return str[2:]
	}

	return ""
}
