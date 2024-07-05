package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {
	cards := deck.New(deck.WithSort(deck.SortBySuit))
	for _, card := range cards {
		fmt.Println(card)
	}
}
