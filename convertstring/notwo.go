package main

import (
	"fmt"
	"strconv"
)

func stringToNumberArray(str string) []int {
	result := make([]int, 0)
	for _, char := range str {
		if num, err := strconv.Atoi(string(char)); err == nil {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	inputString := "1AB23C5678D"
	resultArrayB := stringToNumberArray(inputString)
	fmt.Println(resultArrayB) // Output: [1 2 3 5 6 7 8]
}
