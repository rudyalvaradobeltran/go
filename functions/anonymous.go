package functions

import "fmt"

func Operation() {
	var number3 int = 32
	var number4 int = 243
	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}
	sum = func(number1 int, number2 int) int {
		return number1 + number2 + number3 + number4
	}
	fmt.Println(sum(10, 25))
}
