package main

import "fmt"

func main() {

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println(hand.toString())

}
