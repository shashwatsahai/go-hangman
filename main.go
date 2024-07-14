package main

import (
	"fmt"

	"github.com/shashwatsahai/go-hangman/game"
)

func main() {
	g1 := game.NewGame("john", 4)
	var guess rune
	var win bool
	for g1.TriesLeft > 0 && !win {
		current := g1.CurrentState()
		if current == g1.Word {
			win = true
			fmt.Println("You Won")
			continue
		}
		fmt.Println(g1.CurrentState())
		fmt.Println("Tries Left", g1.TriesLeft)

		fmt.Println("INPUT YOUR GUESS")
		fmt.Scanf("%c\n", &guess)

		g1.Guess(guess)

	}
}
