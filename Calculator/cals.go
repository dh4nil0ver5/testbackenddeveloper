package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculator(input string) bool {
	// Replace ":" with "/" for division
	input = strings.Replace(input, ":", "/", -1)

	// Split the input string by operators
	operands := strings.FieldsFunc(input, func(r rune) bool {
		return r == '+' || r == '-' || r == '*' || r == '/'
	})

	// Initialize result variable
	result := 0

	// Iterate over operands and perform arithmetic operations
	for i, operand := range operands {
		// Parse the operand to integer
		num, err := strconv.Atoi(operand)
		if err != nil {
			fmt.Println("Error parsing operand:", err)
			return false
		}

		// Calculate modulus if it's not the last operand
		if i < len(operands)-1 {
			if num%5 != 0 {
				return false
			}
		}

		// Perform addition, subtraction, multiplication, or division based on the operator
		if i > 0 {
			operatorIndex := strings.IndexAny(input, "+-*/")
			if operatorIndex == -1 {
				fmt.Println("Invalid input: operator not found")
				return false
			}
			operator := input[operatorIndex]
			switch operator {
			case '+':
				result += num
			case '-':
				result -= num
			case '*':
				result *= num
			case '/':
				if num == 0 {
					fmt.Println("Division by zero error")
					return false
				}
				result /= num
			}
			input = input[operatorIndex+1:]
		} else {
			result = num
			input = input[len(operand):]
		}
	}

	return result == 15
}

func main() {
	input := "5+5+5*5/5"
	fmt.Println("Result is equal to 5:", calculator(input))
}
