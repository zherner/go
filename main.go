package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	welcomeMsg = "Starting ++=_P|nky-Pra[t|{e_=++"
	goodbyeMsg = "Goodbye!"
)

// checkUserInput checks user input returns true if strIn matches strNeed
func checkUserInput(strIn, strNeed string) bool {
	clearScreen()
	if strIn == strNeed {
		fmt.Print("Good job!")
		return true
	}

	fmt.Print("Whomp, whomp...")
	return false
}

// score displays the user's correct attempts out of total turns
func score(p, t int) {
	fmt.Printf("(%d/%d)\n", p, t)
}

// clearScreen clears the screen
func clearScreen() {
	fmt.Print("\033[2J")
}

// moveTopLeftScreen moves the cursor to the top left position of the screen
func moveTopLeftScreen() {
	fmt.Print("\033[H")
}

func main() {

	fmt.Printf("%q\n\n", welcomeMsg)

	pinkyChars := []string{":", "{", "}", "\"", "=", "-", "|", "+", "'", "/", "_", "\\", "p", "P", "[", "]"}

	var userInput string
	points := 0

	for turn := 1; userInput != "exit"; turn++ {
		//pick random element in slice
		rand.Seed(time.Now().UnixNano())
		neededInput := pinkyChars[rand.Intn(len(pinkyChars))]

		fmt.Printf("Type the following: %s\n", neededInput)
		var userInput = ""
		fmt.Scan(&userInput)

		if userInput == "exit" {
			fmt.Printf("%s\n", goodbyeMsg)
			break
		}
		if checkUserInput(userInput, neededInput) {
			points++
		}
		score(points, turn)
	}
}
