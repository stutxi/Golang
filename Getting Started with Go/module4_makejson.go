// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

// Submit your source code for the program,
// “makejson.go”.

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Print("Enter a name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Print("Enter an address: ")
	var address string
	fmt.Scanln(&address)
	userInfo := map[string]string{
		"name":    name,
		"address": address,
	}
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON Object:")
	fmt.Println(string(jsonData))
}
