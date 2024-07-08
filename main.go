package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {
	cards := deck.New(deck.WithMultipleDecks(2), deck.WithSort(deck.SortByValue))
	for _, card := range cards {
		fmt.Println(card)
	}
}
