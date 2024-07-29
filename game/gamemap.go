package game

import (
	"fmt"
	"sync"
)

type GameMap struct {
	mu    sync.Mutex
	games map[string]*Game
}

// type map[int]*game.Game;

func (gm *GameMap) AddGame(id string, g *Game) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	fmt.Print("add", id)
	gm.games[id] = g
}

func (gm *GameMap) GetGame(id string) (*Game, bool) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	fmt.Println(gm.games)
	g, ok := gm.games[id]
	fmt.Println("ID=", id)
	fmt.Println("HERE", g, ok, gm, id)
	return g, ok
}

func NewGameMap() *GameMap {
	return &GameMap{
		games: make(map[string]*Game),
	}
}
