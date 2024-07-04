package deck

import (
	"fmt"
	"log"
	_ "math/rand"
)

type Suit uint8  // Represents the suit of a card
type Value uint8 // Represents the value of a card

const (
	Spades   Suit = iota // value 0
	Diamonds             // value 1
	Clubs                // value 2
	Hearts               // value 3
)

const (
	Ace   Value = iota + 1 // value 1
	Two                    // value 2
	Three                  // value 3
	Four                   // value 4
	Five                   // value 5
	Six                    // value 6
	Seven                  // value 7
	Eight                  // value 8
	Nine                   // value 9
	Ten                    // value 10
	Jack                   // value 11
	Queen                  // value 12
	King                   // value 13
)

// Maps for the string representation of the suit and value
// Allows for readable output when printing a card, e.g. "Ace of Spades" instead of {0, 0}
var suits = map[Suit]string{
	0: "Spades",
	1: "Diamonds",
	2: "Clubs",
	3: "Hearts",
}

var values = map[Value]string{
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

type Card struct {
	Suit  Suit
	Value Value
}

// String() is a special method that's called whenever a print function is used
func (c Card) String() string {
	if c.Value < Ace || c.Value > King {
		return fmt.Sprintf("invalid card value: %d", c.Value)
	}
	if c.Suit < Spades || c.Suit > Hearts {
		return fmt.Sprintf("invalid card suit: %d", c.Suit)
	}
	return fmt.Sprintf("%s of %s", values[c.Value], suits[c.Suit])
}

// Configuration Options
type DeckOptions struct {
	shuffle bool
}

type DeckOptionsFunc func(deckOpts *DeckOptions)

func WithShuffle() DeckOptionsFunc {
	return func(deckOpts *DeckOptions) {
		deckOpts.shuffle = true
	}
}

// Creates a complete deck of cards
func New(opts ...DeckOptionsFunc) []Card {
	defaultConfig := &DeckOptions{
		shuffle: false,
	}

	for _, opt := range opts {
		opt(defaultConfig)
	}

	deckOfCards := []Card{}
	for suit := Spades; suit <= Hearts; suit++ {
		for value := Ace; value <= King; value++ {
			deckOfCards = append(deckOfCards, Card{suit, value})
		}
	}
	log.Println("Successfully created deck of cards")
	return deckOfCards
}
