package main

import (
  "fmt"
  "math/rand"
  "time"
)

const welcomeMsg = "Starting =_P|nky-Pra[t|{e_="

func main() {
	goodbyeMsg := "Goodbye! \n\n"
	pinkyChars := []string{":", "{", "}", "\"", "=", "-", "|", "+", "'", "/", "_", "\\", "p", "P", "[", "]"}

	//pick random element in slice
	rand.Seed(time.Now().UnixNano())
	rand.Intn(len(pinkyChars))

	fmt.Printf(welcomeMsg)

	var userInput string
	//loop here
	for userInput != "exit" {
		neededInput := pinkyChars[rand.Intn(len(pinkyChars))]
		fmt.Println("Type the following:", neededInput)
		var userInput = ""
		fmt.Scan(&userInput)

		if userInput == "exit" {
			fmt.Printf(goodbyeMsg)
			break
		}
		checkUserInput(userInput, neededInput)
	}
