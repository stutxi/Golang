// Write a Bubble Sort program in Go. The program
// should prompt the user to type in a sequence of up to 10 integers. The program
// should print the integers out on one line, in sorted order, from least to
// greatest. Use your favorite search tool to find a description of how the bubble
// sort algorithm works.

// As part of this program, you should write a
// function called BubbleSort() which
// takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
// order.

// A recurring operation in the bubble sort algorithm is
// the Swap operation which swaps the position of two adjacent elements in the
// slice. You should write a Swap() function which performs this operation. Your Swap()
// function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice. The Swap() function should return nothing, but it should swap
// the contents of the slice in position i with the contents in position i+1.

// Submit your Go program source code.

package main

import (
	"fmt"
	"strconv"
)

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}

func main() {
	var input string
	arr := make([]int, 0, 10)

	fmt.Println("Enter up to 10 integers:")

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter integer %d: ", i+1)
		fmt.Scan(&input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			i--
			continue
		}

		arr = append(arr, num)
	}

	fmt.Printf("Unsorted array: %v\n", arr)
	BubbleSort(arr)
	fmt.Printf("Sorted array: %v\n", arr)
}
