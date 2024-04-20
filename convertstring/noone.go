package main

import (
	"fmt"
	"strings"
)

func stringToArray(str string) []string {
	return strings.Split(str, "")
}

func main() {
	inputString := "1AB23C5678D"
	resultArrayA := stringToArray(inputString)
	fmt.Println(resultArrayA) // Output: ["1", "A", "B", "2", "3", "C", "5", "6", "7", "8", "D"]
}
