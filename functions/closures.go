package functions

import "fmt"

func multiplication(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	multiplying := 2
	multiplicationResult := multiplication(multiplying)
	for i := 1; i <= 10; i++ {
		fmt.Println(multiplicationResult())
	}
}
