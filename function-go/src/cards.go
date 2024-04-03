package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Card struct {
	Suit  string
	Value string
}

type Cards struct {
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
        arr = append(arr, string(cards[i].Suit[0:1]+cards[i].Value[0:1]))
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var cardsResponse Cards
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
