package models

type Ability int8

const (
	Nothing Ability = iota
	Peek
	Spy
	Swap
)

var EmptyCard = Card{-1, Nothing}

type Card struct {
	Value   int8
	Ability Ability
}
