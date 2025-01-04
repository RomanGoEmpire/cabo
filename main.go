package main

import (
	"fmt"
	"log"

	"github.com/RomanGoEmpire/cabo/models"
)

func main() {
	player := models.NewPlayer("bob")
	fmt.Println(player)

	deck := models.NewDeck()

	card := deck.DrawCard()

	player.ReplaceCard(card, 0)
	player.ReplaceCard(card, 1)
	fmt.Println(player.Cards)

	removed, err := player.RemoveCards(13)
	if err != nil {
		log.Fatalf("Could not remove card")
	}
	fmt.Println(removed)
}
