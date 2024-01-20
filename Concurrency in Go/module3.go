// Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

// The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Printf("Enter %d integers separated by spaces: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Input array:", array)

	var wg sync.WaitGroup

	partitionSize := n / 4
	parts := make([][]int, 4)

	for i := 0; i < 4; i++ {
		startIndex := i * partitionSize
		endIndex := (i + 1) * partitionSize

		if i == 3 {
			endIndex = n
		}

		parts[i] = array[startIndex:endIndex]

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sort.Ints(parts[i])
			fmt.Printf("Goroutine %d sorted: %v\n", i, parts[i])
		}(i)
	}

	wg.Wait()

	sortedArray := mergeSortedArrays(parts[0], parts[1], parts[2], parts[3])

	fmt.Println("Sorted array:", sortedArray)
}

func mergeSortedArrays(arr1, arr2, arr3, arr4 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2)+len(arr3)+len(arr4))

	var wg sync.WaitGroup

	merge := func(arr []int) {
		defer wg.Done()
		merged = append(merged, arr...)
	}

	wg.Add(4)
	go merge(arr1)
	go merge(arr2)
	go merge(arr3)
	go merge(arr4)

	wg.Wait()

	sort.Ints(merged)
	return merged
}
