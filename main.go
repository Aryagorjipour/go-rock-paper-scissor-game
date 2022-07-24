package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	fmt.Println("Hello There! Let's Play rock paper scissor.")
	var isClose bool = false
	for isClose == false {
		println(playGame())
	}
}

func playGame() string {
	turnChoice := rand.Intn(3)
	var userChoice int
	fmt.Print("Chose you're (1 for rock, 2 for paper, 3 for scissor): ")
	_, error := fmt.Scanln(&userChoice)
	if error != nil {
		log.Fatal(error)
	}
	if turnChoice == 0 { //0 for rock
		switch userChoice {
		case 1:
			return "draw!"
		case 2:
			return "user win"
		case 3:
			return "computer win"
		default:
			fmt.Println("Please choose correct form")
			playGame()
		}
	} else if turnChoice == 1 { //1 for paper
		switch userChoice {
		case 1:
			return "computer win"
		case 2:
			return "draw!"
		case 3:
			return "user win"
		default:
			fmt.Println("Please choose correct form")
			playGame()
		}
	} else if turnChoice == 2 { //2 for scissor
		switch userChoice {
		case 1:
			return "user win"
		case 2:
			return "computer win"
		case 3:
			return "draw!"
		default:
			fmt.Println("Please choose correct form")
			playGame()
		}
	} else {
		fmt.Println("Please choose correct form")
		playGame()
	}
	return "Nothing Happened"
}
