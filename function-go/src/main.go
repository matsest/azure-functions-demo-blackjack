package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Game struct {
	Players []Player
	Winner  string
}

func PlayBlackJack(name1 string, name2 string, cardDeck []Card) Game {

	var cards []Card = cardDeck

	p1 := newPlayer(name1)
	p1.drawCards(&cards, 2)

	p2 := newPlayer(name2)
	p2.drawCards(&cards, 2)

	var winner string

	game := Game{
		Players: []Player{p1, p2},
		Winner:  winner,
	}

	// initialize pointers to players
	player1 := &game.Players[0]
	player2 := &game.Players[1]

	for len(winner) == 0 {

		// Initial Blackjack checks
		if player1.Value == 21 && player2.Value == 21 {
			winner = "Draw"
			break
		}

		if player1.Value == 21 {
			winner = player1.Name
			break
		}

		if player2.Value == 21 {
			winner = player2.Name
			break
		}

		// Player 1 draw cards
		for player1.Value < 17 {
			player1.drawCards(&cards, 1)
		}

		if player1.Value > 21 {
			winner = player2.Name
			break
		}

		// Player 2 draw cards
		for player2.Value < player1.Value {
			player2.drawCards(&cards, 1)
		}

		// Check winners
		if player2.Value == player1.Value {
			winner = "Draw"
		} else if player2.Value > 21 {
			winner = player1.Name
		} else {
			winner = player2.Name
		}
	}

	game.Winner = winner
	return game
}

func play(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	cards := downloadCards("http://nav-deckofcards.herokuapp.com/shuffle")

	fmt.Fprint(w, "cards: ", abbreviateCards(cards[:]...), "\n")
	fmt.Fprint(w, "points: ", cardsValues(cards[:]...), "\n")

	game := PlayBlackJack("Mats", "Line", cards)
	for i := 0; i < len(game.Players); i++ {
		fmt.Fprintf(w, "%s: %v Sum: %d \n", game.Players[i].Name, abbreviateCards(game.Players[i].Hand...), game.Players[i].Value)
	}
	fmt.Fprint(w, "Winner: ", game.Winner, "\n")

	duration := time.Since(start)

	log.Println(r.Method, r.RequestURI, r.UserAgent(), duration)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/blackjack", play)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
