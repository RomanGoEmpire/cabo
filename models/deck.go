package models

type Deck struct {
	Draw    []*Card
	Deposit []*Card
}

func (d *Deck) DrawCard() *Card {
	lastIndex := len(d.Draw) - 1
	lastElement := d.Draw[lastIndex]
	d.Draw = d.Draw[:lastIndex]
	return lastElement
}

func NewDeck() Deck {
	cards := []*Card{}

	cardAmounts := []struct {
		Value   int
		Ability Ability
		Amount  int
	}{
		{0, Nothing, 2},
		{1, Nothing, 4},
		{2, Nothing, 4},
		{3, Nothing, 4},
		{4, Nothing, 4},
		{5, Nothing, 4},
		{6, Nothing, 4},
		{7, Peek, 4},
		{8, Peek, 4},
		{9, Spy, 4},
		{10, Spy, 4},
		{11, Swap, 4},
		{12, Swap, 4},
		{13, Nothing, 2},
	}

	for _, cardType := range cardAmounts {
		for i := 0; i < cardType.Amount; i++ {
			cards = append(cards, &Card{int8(cardType.Value), cardType.Ability})
		}

	}
	// rand.Shuffle(len(cards), func(i int, j int) {
	// 	cards[i], cards[j] = cards[j], cards[i]
	// })

	return Deck{cards, []*Card{}}
}
