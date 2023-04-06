package main

func main() {
	cards := newDeckFromFile("my_cards")

	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// fmt.Println("separate")
	// remainingCards.print()

	// cards.saveToFile("my_cards")

	cards.shuffle()
	cards.print()
}
