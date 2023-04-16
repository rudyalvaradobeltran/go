package main

import (
	"fmt"

	"github.com/rudyalvaradobeltran/go/variables"
)

func main() {
	variables.ShowIntegers()
	variables.ShowRestOfVariables()
	state, text := variables.ConvertToText(666)
	fmt.Println("State = ", state)
	fmt.Println("Text = ", text)
}
