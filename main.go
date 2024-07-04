package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {
	cards := deck.New()
	for _, card := range cards {
		fmt.Println(card)
	}
}
