package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var text string
var err error

func EnterNumbers() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter number 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error: " + err.Error())
		}
	}
	fmt.Println("Enter number 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error: " + err.Error())
		}
	}
	fmt.Println("Enter text: ")
	if scanner.Scan() {
		text = scanner.Text()
		if err != nil {
			panic("Error: " + err.Error())
		}
	}
	fmt.Println(text, number1*number2)
}

func MultiplicationConcat() string {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a number")
	if scanner.Scan() {
		number, error := strconv.Atoi(scanner.Text())
		if error != nil {
			fmt.Println("Invalid number")
			MultiplicationConcat()
			return text
		}
		for i := 1; i <= 10; i++ {
			text += fmt.Sprintf("%d x %d: %d \n", i, number, i*number)
		}
	}
	return text
}
