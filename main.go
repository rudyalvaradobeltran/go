package main

import (
	"fmt"
	"runtime"

	"github.com/rudyalvaradobeltran/go/exercises"
	"github.com/rudyalvaradobeltran/go/input"
	"github.com/rudyalvaradobeltran/go/iterations"
	"github.com/rudyalvaradobeltran/go/variables"
)

func main() {
	// integers
	variables.ShowIntegers()

	// rest of variables and convert to text
	variables.ShowRestOfVariables()
	state, text := variables.ConvertToText(666)
	fmt.Println("State = ", state)
	fmt.Println("Text = ", text)

	// conditions
	if os := runtime.GOOS; os == "linux" || os == "darwin" {
		fmt.Println("Not windows")
	} else {
		fmt.Println("Windows")
	}
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("This is %s \n", os)
	}

	// exercise 01
	fmt.Println(exercises.GreaterThan("99"))
	fmt.Println(exercises.GreaterThan("101"))

	// stdin example
	input.EnterNumbers()

	// iterations
	iterations.Iterate()

	// exercise 02
	exercises.Multiplication()
}
