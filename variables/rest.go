package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func ShowRestOfVariables() {
	Name = "Peter"
	State = true
	Salary = 1900000.99
	Date = time.Now()
	fmt.Println("Name = ", Name)
	fmt.Println("State = ", State)
	fmt.Println("Salary = ", Salary)
	fmt.Println("Date = ", Date)
}

func ConvertToText(number int) (bool, string) {
	return true, strconv.Itoa(number)
}
