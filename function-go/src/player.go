package main

type Player struct {
	Name  string
	Hand  []Card
	Value int
}

func (player *Player) drawCards(cards *[]Card, number int) {
	newCards, remainingCards := drawCardsFromTop(*cards, number)
	player.Hand = append(player.Hand, newCards...)
	*cards = remainingCards
	player.Value = cardsValues(player.Hand...)
}

func newPlayer(name string) Player {
	p := Player{}
	p.Name = name
	p.Hand = []Card{}
	p.Value = 0
	return p
}
