package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rudyalvaradobeltran/go/input"
)

var fileName string = "./files/txt/file.txt"

func SaveMultiplicationToFile() {
	var text string = input.MultiplicationConcat()
	file, error := os.Create(fileName)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddMultiplicationToFile() {
	var text string = input.MultiplicationConcat()
	if !Append(text) {
		fmt.Println("Error: Cannot concatenate text")
	}
}

func Append(text string) bool {
	file, error := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if error != nil {
		fmt.Println("Error on Append: " + error.Error())
		return false
	}
	_, error = file.WriteString(text)
	if error != nil {
		fmt.Println("Error on WriteString: " + error.Error())
		return false
	}
	file.Close()
	return true
}

func ReadFileButDeprecated() {
	file, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error on Reading file: " + error.Error())
		return
	}
	fmt.Println(string(file))
}

func ReadFile() {
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error on Reading file: " + error.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		record := scanner.Text()
		fmt.Println("> " + record)
	}
	file.Close()
}
