package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Card struct {
	Suit  string
	Value string
}

type CardsResp struct {
	Cards []Card
}

func cardValue(card Card) int {
	switch card.Value {
	case "JACK", "QUEEN", "KING":
		return 10
	case "ACE":
		return 11
	}
	res, err := strconv.Atoi(card.Value)
	if err != nil {
		log.Fatalln(err)
	}

	return res
}

func abbreviateCards(cards ...Card) []string {
	arr := []string{}
	for i := 0; i < len(cards); i++ {
		var valToAdd string
		if len(cards[i].Value) > 2 {
			valToAdd = string(cards[i].Value[0])
		} else {
			valToAdd = string(cards[i].Value)
		}
		arr = append(arr, string(cards[i].Suit[0:1])+valToAdd)
	}
	return arr
}

func cardsValues(cards ...Card) int {

	length := len(cards)
	sum := 0
	for i := 0; i < length; i++ {
		sum += cardValue(cards[i])
	}
	return sum
}

func downloadCards(uri string) []Card {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var cardsResponse CardsResp
	err = json.Unmarshal([]byte(body), &cardsResponse)
	if err != nil {
		log.Fatalln(err)
	}

	var cards []Card
	for _, card := range cardsResponse.Cards {
		cards = append(cards, Card{
			Value: card.Value,
			Suit:  card.Suit,
		})
	}

	return cards
}

func drawCardsFromTop(cards []Card, number int) ([]Card, []Card) {
	res1 := cards[0:number]
	res2 := cards[number:]
	return res1, res2
}
