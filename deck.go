//go:generate stringer -type=Rank,Suit

// Package deck - for a deck of cards
package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank Rank = Ace
	maxRank Rank = King
)

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return c.Rank.String() + " of " + c.Suit.String() + "s"
}

func newDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absValue(cards[i]) < absValue(cards[j])
	}
}

func absValue(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func shufflePerm(cards []Card) []Card {
	ret := make([]Card, len(cards))

	rand.Seed(time.Now().UnixNano())
	indicies := rand.Perm(len(cards))

	for i, card := range cards {
		ret[i] = card[indicies[i]]
	}
	return ret
}

func shuffle(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}
