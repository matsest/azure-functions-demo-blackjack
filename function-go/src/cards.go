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

func cardValue(card Card) int {
	switch card.Value {
	case "J", "Q", "K":
		return 10
	case "A":
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
		arr = append(arr, string(cards[i].Suit[0:1]+cards[i].Value))
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

	var cards []Card
	err = json.Unmarshal([]byte(body), &cards)
	if err != nil {
		log.Fatalln(err)
	}

	return cards
}

func drawCardsFromTop(cards []Card, number int) ([]Card, []Card) {
	res1 := cards[0:number]
	res2 := cards[number:(len(cards))]
	return res1, res2
}
