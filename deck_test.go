package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Ten, Suit: Heart})
	fmt.Println(Card{Rank: Queen, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Spades
	// Ten of Hearts
	// Queen of Diamonds
	// Joker
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Errorf("First card is not Ace fo Spaces: got %v", cards[0])
	}
}
func TestShuffle(t *testing.T) {
	cards := newDeck()
	dontWant := cards[0]

	shuffle(cards)
	got := cards[0]
	if got != dontWant {
		t.Errorf("Expected %v, got %v", got, dontWant)
	}
}
func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}
func TestString(t *testing.T) {

	c := Card{Ace, Spade}
	expected := "Ace of Spades"
	got := c.String()
	if got != expected {
		t.Errorf("String : expected %s, got %s", expected, got)
	}
	c = Card{Suit: Joker}
	expected = "Joker"
	got = c.String()
	if got != expected {
		t.Errorf("String : expected %s, got %s", expected, got)
	}
}
