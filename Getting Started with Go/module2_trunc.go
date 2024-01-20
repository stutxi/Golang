// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.

// Submit your source code for the program, “trunc.go”.

package main

import "fmt"

func main() {
	var input float64

	fmt.Println("Enter a floating point number: ")
	fmt.Scan(&input)

	truncatedInt := int(input)

	fmt.Println("Truncated Integer: ", truncatedInt)
}
