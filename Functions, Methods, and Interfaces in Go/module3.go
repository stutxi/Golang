// Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

// Animal  |  Food eaten  |  Locomotion method
// cow     |  grass       |  walk
// bird    |  worms       |  fly
// snake   |  mice        |  slither
//         |              |

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

// You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.

// Submit your Go program source code.

package main

import "fmt"

type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.Food)
}
func (a Animal) Move() {
	fmt.Println(a.Locomotion)
}
func (a Animal) Speak() {
	fmt.Println(a.Noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   {Food: "grass", Locomotion: "walk", Noise: "moo"},
		"bird":  {Food: "worms", Locomotion: "fly", Noise: "peep"},
		"snake": {Food: "mice", Locomotion: "slither", Noise: "hsss"},
	}

	for {
		fmt.Print(">")
		var animalName, infoType string
		fmt.Scan(&animalName, &infoType)

		animal, exists := animals[animalName]

		if !exists {
			fmt.Println("Invalid animal name. Please choose cow, bird, or snake.")
			continue
		}

		switch infoType {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid information type. Please choose eat, move, or speak.")
		}
	}
}
