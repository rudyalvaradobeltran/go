package main

import (
	"fmt"
	"runtime"

	"github.com/rudyalvaradobeltran/go/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowRestOfVariables()
	state, text := variables.ConvertToText(666)
	fmt.Println("State = ", state)
	fmt.Println("Text = ", text)
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
}
