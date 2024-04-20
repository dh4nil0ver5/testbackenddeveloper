package main

import (
	"fmt"
)

// visualizeArray function displays a numerical array as vertical bar charts
func visualizeArray(arr []int) {
	// Find the maximum value in the array
	maxValue := 0
	for _, num := range arr {
		if num > maxValue {
			maxValue = num
		}
	}

	// Display the vertical bar charts
	for i := maxValue; i > 0; i-- {
		for _, num := range arr {
			if num >= i {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	// Display the horizontal axis (array values)
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

// insertionSort function implements the insertion sort algorithm and visualizes each step
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key

		// Visualize each step of the sorting process
		visualizeArray(arr)
	}
}

func main() {
	// Input numerical array
	arr := []int{1, 4, 5, 6, 8, 2}

	// Display the original array as vertical bar charts
	fmt.Println("Original array:")
	visualizeArray(arr)
	fmt.Println()

	// Sort the array using insertion sort and visualize each step
	fmt.Println("Sorting steps:")
	insertionSort(arr)
}