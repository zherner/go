package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	welcomeMsg = "Starting ++=_P|nky-Pra[t|{e_=++"
	goodbyeMsg = "Goodbye!"
)

// checkUserInput checks user input returns true if strIn matches strNeed
func checkUserInput(strIn, strNeed rune) bool {
	clearScreen()
	fmt.Println("++=_P|nky-Pra[t|{e_=++")
	fmt.Printf("Press ESC to quit\n\n")
	if strIn == strNeed {
		fmt.Println("Good job!")
		return true
	}

	fmt.Println("Whomp, whomp...")
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
	// welcome message
	fmt.Printf("%q\n\n", welcomeMsg)
	// initialize points
	points := 0

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	pinkyRunes := []rune{':', '{', '}', '"', '=', '-', '|', '+', '\'', '/', '_', '\\', 'p', 'P', '[', ']'}

	fmt.Println("Press ESC to quit")
	for turn := 1; ; turn++ {
		// pick random element in slice
		rand.Seed(time.Now().UnixNano())
		neededInput := pinkyRunes[rand.Intn(len(pinkyRunes))]

		fmt.Printf("Type the following: %s\n", string(neededInput))

		// get input immediately
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		// handle input
		switch {
		// bail if ESC
		case key == keyboard.KeyEsc:
			fmt.Printf("%s\n", goodbyeMsg)
			return
		case checkUserInput(char, neededInput):
			points++
			score(points, turn)
		default:
			score(points, turn)
		}
	}

}
