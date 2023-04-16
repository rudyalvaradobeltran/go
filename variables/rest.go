package variables

import (
	"fmt"
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
