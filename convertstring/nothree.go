package main

import (
	"fmt"
	"unicode"
)

func stringToCharacterArray(str string) []string {
	result := make([]string, 0)
	for _, char := range str {
		if !unicode.IsDigit(char) {
			result = append(result, string(char))
		}
	}
	return result
}

func main() {
	inputString := "1AB23C5678D"
	resultArrayC := stringToCharacterArray(inputString)
	fmt.Println(resultArrayC) // Output: ["A" "B" "C" "D"]
}
