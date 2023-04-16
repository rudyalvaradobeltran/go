package iterations

import (
	"fmt"
)

func Iterate() {
	a := 1
	for a <= 5 {
		fmt.Println("A: ", a)
		a++
	}
	for b := 3; b <= 9; b += 3 {
		fmt.Println("B: ", b)
	}
	for c := 8; c >= 0; c -= 2 {
		fmt.Println("C: ", c)
	}
	for d := 5; d > 0; d-- {
		fmt.Println("D: ", d)
	}
	for e := 1; e <= 10; e++ {
		if e == 6 {
			break
		}
		fmt.Println("E: ", e)
	}
	for f := 1; f <= 3; f++ {
		if f == 2 {
			continue
		}
		fmt.Println("F: ", f)
	}
}
