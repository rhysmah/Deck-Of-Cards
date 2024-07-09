package deck

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Suit uint8  // Represents the suit of a card
type Value uint8 // Represents the value of a card

// Represents a playing card
type Card struct {
	Suit  Suit
	Value Value
}

// Sorts the deck of cards by suit
type BySuit []Card

func (a BySuit) Len() int           { return len(a) }
func (a BySuit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuit) Less(i, j int) bool { return a[i].Suit < a[j].Suit }

func SortBySuit(deckOfCards []Card) {
	sort.Sort(BySuit(deckOfCards))
}

// Sorts the deck of cards by value
type ByValue []Card

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

func SortByValue(deckOfCards []Card) {
	sort.Sort(ByValue(deckOfCards))
}

const (
	Spades    Suit = iota // value 0
	Diamonds              // value 1
	Clubs                 // value 2
	Hearts                // value 3
	JokerSuit             // value 4
)

const (
	JokerValue Value = iota // value 0
	Ace                     // value 1
	Two                     // value 2
	Three                   // value 3
	Four                    // value 4
	Five                    // value 5
	Six                     // value 6
	Seven                   // value 7
	Eight                   // value 8
	Nine                    // value 9
	Ten                     // value 10
	Jack                    // value 11
	Queen                   // value 12
	King                    // value 13
)

// Maps for the string representation of the suit and value
// Allows for readable output when printing a card, e.g. "Ace of Spades" instead of {0, 0}
var suits = map[Suit]string{
	0: "Spades",
	1: "Diamonds",
	2: "Clubs",
	3: "Hearts",
	4: "Joker",
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

// String() is a special method that's called whenever a print function is used
func (c Card) String() string {
	if c.Suit == JokerSuit {
		return fmt.Sprint(suits[c.Suit])
	}
	if c.Value < Ace || c.Value > King {
		return fmt.Sprintf("invalid card value: %d", c.Value)
	}
	if c.Suit < Spades || c.Suit > Hearts {
		return fmt.Sprintf("invalid card suit: %d", c.Suit)
	}
	return fmt.Sprintf("%s of %s", values[c.Value], suits[c.Suit])
}

// Configuration options for a deck of cards
type DeckOptions struct {
	sortFunc   func([]Card)
	shuffle    bool
	filterCard func(card Card) bool
	numJokers  int
	numDecks   int
}

// Function that takes in a pointer to a DeckOptions struct and modifies it
// This allows for the creation of a deck of cards with different configurations
type DeckOptionsFunc func(deckOpts *DeckOptions)

// Determines how a deck of cards should be sorted
func WithSort(sortFunc func([]Card)) DeckOptionsFunc {
	return func(DeckOpts *DeckOptions) {
		DeckOpts.sortFunc = sortFunc
	}
}

// Determines if a deck of cards should be shuffled
func WithShuffle() DeckOptionsFunc {
	return func(deckOpts *DeckOptions) {
		deckOpts.shuffle = true
	}
}

// Determines if a deck of cards should be filtered
// Decks can be filtered by suit, value, or both
func WithFilteredCards(filterFunc func(card Card) bool) DeckOptionsFunc {
	return func(deckOpts *DeckOptions) {
		deckOpts.filterCard = filterFunc
	}
}

// Determines how many jokers should be added to a deck of cards
func WithJokers(n int) DeckOptionsFunc {
	return func(deckOpts *DeckOptions) {
		deckOpts.numJokers = n
	}
}

// Determines how many decks should be combined to create a deck of cards
func WithMultipleDecks(n int) DeckOptionsFunc {
	return func(deckOpts *DeckOptions) {
		deckOpts.numDecks = n
	}
}

// Creates a complete deck of cards
func New(opts ...DeckOptionsFunc) []Card {
	defaultConfig := &DeckOptions{
		sortFunc:   nil,
		shuffle:    false,
		filterCard: nil,
		numJokers:  0,
		numDecks:   1,
	}

	for _, opt := range opts {
		opt(defaultConfig)
	}

	// Create standard deck of cards
	deckOfCards := []Card{}
	for suit := Spades; suit <= Hearts; suit++ {
		for value := Ace; value <= King; value++ {
			card := Card{Suit: suit, Value: value}

			// Filter cards (if applicable)
			if defaultConfig.filterCard == nil || !defaultConfig.filterCard(card) {
				deckOfCards = append(deckOfCards, card)
			}
		}
	}

	// Combine multiple decks (if applicable)
	if defaultConfig.numDecks > 1 {
		combinedDeck := make([]Card, 0, len(deckOfCards)*defaultConfig.numDecks)
		for i := 0; i < defaultConfig.numDecks; i++ {
			combinedDeck = append(combinedDeck, deckOfCards...)
		}
		deckOfCards = combinedDeck
	}

	// Add Jokers (if applicable)
	for i := 0; i < defaultConfig.numJokers; i++ {
		deckOfCards = append(deckOfCards, Card{Suit: JokerSuit, Value: JokerValue})
	}

	// Shuffle cards (if applicable)
	if defaultConfig.shuffle {
		shuffle(deckOfCards)
	}

	if defaultConfig.sortFunc != nil {
		defaultConfig.sortFunc(deckOfCards)
	}

	log.Println("Successfully created deck of cards")
	return deckOfCards
}

func shuffle(deckOfCards []Card) {
	rand.Shuffle(len(deckOfCards), func(i, j int) {
		deckOfCards[i], deckOfCards[j] = deckOfCards[j], deckOfCards[i]
	})
}
