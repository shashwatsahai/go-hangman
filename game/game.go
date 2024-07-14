package game

import "fmt"

type Game struct {
	Word      string
	Guesses   []rune
	GuessMap  map[rune]bool
	MaxTries  int
	TriesLeft int
}

func NewGame(word string, maxtries int) *Game {
	if word != "" && maxtries > 0 {
		// var game Game
		game := Game{
			Word:      word,
			Guesses:   []rune{},
			GuessMap:  map[rune]bool{},
			MaxTries:  maxtries,
			TriesLeft: maxtries,
		}
		return &game
	}
	panic("INVALID VALUES")
}

func (g *Game) Guess(alphabet rune) bool {
	fmt.Println("INSIDE GUESS")
	for i := 0; i < len(g.Word); i++ {
		if rune(g.Word[i]) == alphabet {
			fmt.Printf("MATCHED %c == %c\n", g.Word[i], alphabet)
			g.Guesses = append(g.Guesses, alphabet)
			g.GuessMap[alphabet] = true
			return true
		}
	}

	g.TriesLeft--
	fmt.Printf("MISSED %c\n", alphabet)
	return false
}

func (g *Game) CurrentState() string {
	curr := ""

	for i := 0; i < len(g.Word); i++ {
		word := g.Word[i]
		_, ok := g.GuessMap[rune(word)]
		if ok {
			curr = curr + string(word)
		} else {
			curr = curr + "_"
		}
	}
	return curr
}
