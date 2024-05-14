package cards

import "math/rand"

type CardType int

const (
	Major CardType = iota
	Minor
)

var cardTypeToString = map[CardType]string{
	Major: "Major Arcana",
	Minor: "Minor Arcana",
}

func (ct CardType) String() string {
	if s, ok := cardTypeToString[ct]; ok {
		return s
	}
	return "unknown"
}

type Card struct {
	Name         string
	Description  string
	RDescription string
	Message      string
	Type         CardType
	Url          string
	Image        string
}

type CardDrawn struct {
	Card     Card
	Reversed bool
}

func NewDeck() []string {
	deck := make([]string, 0)
	for s := range cards {
		deck = append(deck, s)
	}
	return deck
}

func SuffleDeck(deck []string) []string {
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck
}

func DrawCards(deck []string, num int) []CardDrawn {
	hand := make([]CardDrawn, 0)
	for i := 0; i < num; i++ {
		randomIndex := rand.Intn(len(deck) - 1)

		card := GetCard(deck[randomIndex])
		reversed := rand.Intn(2) == 1
		drawnCard := CardDrawn{Card: card, Reversed: reversed}

		deck = append(deck[:randomIndex], deck[randomIndex+1:]...)

		hand = append(hand, drawnCard)
	}
	return hand
}

func GetCard(s string) Card {
	if c, ok := cards[s]; ok {
		return c
	}
	return Card{}
}
