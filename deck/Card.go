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

func (c Card) String() (string, error) {
	if c.Value < Ace || c.Value > King {
		return "", fmt.Errorf("invalid card value: %d", c.Value)
	}
	if c.Suit < Spades || c.Suit > Hearts {
		return "", fmt.Errorf("invalid card suit: %d", c.Suit)
	}
	return fmt.Sprintf("%s of %s", values[c.Value], suits[c.Suit]), nil
}

// Configuration Options
type DeckOptions struct {
	shuffle bool // Determines if cards are initially shuffled
	// sortMethod func([]Card)         // A custom way to sort the Deck
	// numJokers  int                  // Determines how many jokers are added to the deck
	// filterFunc func(card Card) bool // A function for removing a particular card
}

type DeckOptionsFunc func(deckopts *DeckOptions)

func WithShuffle() DeckOptionsFunc {
	return func(deckopts *DeckOptions) {
		deckopts.shuffle = true
	}
}

// Creates a complete deck of cards
func New() []Card {
	deckOfCards := []Card{}
	for suit := Spades; suit <= Hearts; suit++ {
		for value := Ace; value <= King; value++ {
			deckOfCards = append(deckOfCards, Card{suit, value})
		}
	}
	log.Println("Successfully created deck of cards")
	return deckOfCards
}
