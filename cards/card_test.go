package cards

import "testing"

func TestShuffle(t *testing.T) {
	deck1 := NewDeck()
	deck2 := make([]string, len(deck1))
	_ = copy(deck2, deck1)

	deck1 = SuffleDeck(deck1)

	if equal(deck1, deck2) {
		t.Error("Deck1 is equal to Deck2!")
	}
}

func equal(d1, d2 []string) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i := range d1 {
		if d1[i] != d2[i] {
			return false
		}
	}
	return true
}
