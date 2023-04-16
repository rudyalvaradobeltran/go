package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiplication() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number")
	if scanner.Scan() {
		number, error := strconv.Atoi(scanner.Text())
		if error != nil {
			fmt.Println("Invalid number")
			Multiplication()
			return
		}
		for i := 1; i <= 10; i++ {
			fmt.Println(i, "x", number, ": ", i*number)
		}
	}
}
