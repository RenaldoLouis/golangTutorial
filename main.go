package main

import (
	decks "golangTutorial/deck"
)

var (
	decksVar decks.InterfaceDecks = decks.NewDecks()
)

func main() {
	cards := decksVar.NewDeckFromFile("my_cards")

	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// fmt.Println("separate")
	// remainingCards.print()

	// cards.saveToFile("my_cards")

	cards.Shuffle()
	cards.Print()

}
