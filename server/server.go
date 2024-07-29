package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/shashwatsahai/go-hangman/game"
)

type ServerStart struct {
	Hosturl string
	Conn    *http.Server
}

var games = game.NewGameMap()

func (s *ServerStart) Start() error {
	s.Conn = &http.Server{
		Addr:    s.Hosturl,
		Handler: http.DefaultServeMux,
	}
	log.Printf("Starting server at %s\n", s.Hosturl)

	// Define a simple handler
	http.HandleFunc("/newgame", func(w http.ResponseWriter, r *http.Request) {
		g1 := game.NewGame("john", 4)
		id := uuid.New().String()
		games.AddGame(id, g1)
		w.Write([]byte(id))
	})

	http.HandleFunc("/guess", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		// fmt.Print(query)
		id := query.Get("id")

		fmt.Print(id)
		// fmt.Print("here")
		var resBody struct {
			Guess string `json:"guess"`
		}

		// var res1 resBody
		err := json.NewDecoder(r.Body).Decode(&resBody)

		if err != nil {
			fmt.Print(err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if len(resBody.Guess) != 1 {
			http.Error(w, "Guess must be a single character", http.StatusBadRequest)
			return
		}

		guess := rune(resBody.Guess[0])

		fmt.Println("REQID=", id)
		gameInstance, ok := games.GetGame(id)
		fmt.Print("OK=", ok, gameInstance)
		if !ok {
			fmt.Print("No game")
			w.Write([]byte("No game found"))
			return
		}

		// if g1
		gameInstance.Guess(guess)
		if gameInstance.IsWon() {
			fmt.Println("You Won")
			w.Write([]byte("You won"))
		}
		if gameInstance.IsLost() {
			fmt.Println("You Lost")
			w.Write([]byte("You Lost"))
		}

		w.Header().Set("Content-Type", "application/json")

		response := struct {
			CurrentWord string `json:"currentWord"`
			TriesLeft   int    `json:"triesLeft"`
			Status      string `json:"status"`
		}{
			CurrentWord: gameInstance.CurrentState(),
			TriesLeft:   gameInstance.TriesLeft,
			Status:      "Guess processed",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Write(jsonResponse)
	})

	err := s.Conn.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", s.Hosturl, err)
		return err
	}
	return nil
}
