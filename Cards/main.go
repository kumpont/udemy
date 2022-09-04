package main

import (
	"fmt"
	"os"
)

func main() {

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println(hand.toString())
	err := hand.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	newdeck, err := newDeckFromFile("my_cards")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	newdeck.print()

	cards.shuffle()
	cards.print()

}
