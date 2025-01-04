package main

import (
	"testing"

	"github.com/RomanGoEmpire/cabo/models"
)

const (
	RED   = "\033[36m"
	RESET = "\033[0m"
)

func FormatErrorMessage(t *testing.T, want interface{}, got interface{}) {
	t.Fatalf("\nWant: %v\nGot : %s%v%s", want, RED, got, RESET)
}

func TestNewDeck(t *testing.T) {
	// Sum of all cards
	want := 338

	deck := models.NewDeck()
	got := 0
	for _, card := range deck.Draw {
		got += int(card.Value)
	}

	if got != want {
		FormatErrorMessage(t, got, want)
	}
}
