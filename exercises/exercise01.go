package exercises

import (
	"strconv"
)

func GreaterThan(text string) (int, string) {
	number, _ := strconv.Atoi(text)
	if number > 100 {
		return number, "is greater than 100"
	} else {
		return number, "is lower or equal than 100"
	}
}
