package main

import (
	"fmt"
	"math/rand"
)

type Ability int

const (
	Nothing = iota
	Peek
	Spy
	Swap
)

type Card struct {
	Value   int8
	Ability Ability
}

type Player struct {
	Name  string
	Score int8
	cards []Card
}

func NewPlayer(name string) Player {
	return Player{
		Name:  name,
		Score: 0,
		cards: nil,
	}
}

func (p *Player) AddCards(cards []Card) {
	p.cards = cards
}

func NewDeck() [52]Card {
	cards := [52]Card{}

	cards[0] = Card{0, Nothing}
	cards[1] = Card{0, Nothing}
	cards[2] = Card{13, Nothing}
	cards[3] = Card{13, Nothing}

	for i := 1; i <= 12; i++ {
		start := i * 4
		for c := 0; c < 4; c++ {
			cardType := Nothing
			switch i {
			case 7, 8:
				cardType = Peek
			case 9, 10:
				cardType = Spy
			case 11, 12:
				cardType = Swap
			}
			cards[start+c] = Card{int8(i), Ability(cardType)}
		}
	}

	rand.Shuffle(len(cards), func(i int, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func main() {
	player := NewPlayer("Bob")
	fmt.Println(player)
	fmt.Println(player.cards, len(player.cards))

	cards := NewDeck()
	fmt.Println(cards)
}
