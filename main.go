package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	welcomeMsg = "++=_P|nky-Pra[t|{e_=++"
	goodbyeMsg = "Goodbye!"
)

var (
	// initialize points
	points int = 0
)

// checkUserInput clears screen,
// checks user input,
// returns true if strIn matches strNeed
func checkUserInput(strIn, strNeed rune) bool {
	clearScreen()
	moveTopLeftScreen()
	fmt.Printf("%q\n\n", welcomeMsg)
	fmt.Printf("Press ESC to quit\n")
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
	// init pinky keys as runes
	pinkyRunes := []rune{':', '{', '}', '"', '=', '-', '|', '+', '\'', '/', '_', '\\', 'p', 'P', '[', ']'}

	// clean screen
	clearScreen()
	moveTopLeftScreen()

	// welcome
	fmt.Printf("%q\n\n", welcomeMsg)
	fmt.Println("Press ESC to quit")

	// get keys
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	// start game loop, counting turns for score
	for turn := 1; ; turn++ {
		// pick random element in slice
		rand.Seed(time.Now().UnixNano())
		neededInput := pinkyRunes[rand.Intn(len(pinkyRunes))]

		// get input immediately
		fmt.Printf("Type the following: %s\n", string(neededInput))
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
