package models

import (
	"log"
)

type Player struct {
	Score int8
	Name  string
	Cards [4]*Card
}

func NewPlayer(name string) Player {
	return Player{
		Name:  name,
		Score: 0,
		Cards: [4]*Card{nil, nil, nil, nil},
	}
}

func (p *Player) ReplaceCard(card *Card, index int) (RemovedCard *Card) {
	if index < 0 || index > 3 {
		log.Fatal("Index must be 0 - 3")
	}
	RemovedCard, p.Cards[index] = p.Cards[index], card
	return
}

// Removes and returns card(s) of passed Value.
func (p *Player) RemoveCards(Value int8) (RemovedCards []*Card, Error error) {
	if Value < 0 || Value > 13 {
		log.Fatal("Value must be 0 - 12")
	}
	for index, card := range p.Cards {

		if card == nil {
			continue
		}

		if card.Value == Value {
			RemovedCards = append(RemovedCards, card)
			p.Cards[index] = nil
		}
	}
	return
}
