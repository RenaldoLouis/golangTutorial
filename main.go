package main

import (
	"fmt"
	decks "golangTutorial/deck"
	"golangTutorial/structs"
)

var (
	decksVar decks.InterfaceDecks = decks.NewDecks()
)

func main() {
	//Part 1 Cards
	cards := decksVar.NewDeckFromFile("my_cards")

	//no need to run this START
	// hand, remainingCards := deal(cards, 2)
	// hand.print()
	// fmt.Println("separate")
	// remainingCards.print()

	// cards.saveToFile("my_cards")
	//no need to run this END

	cards.Shuffle()
	// cards.Print()

	//Part 2 Structs
	// alex := structs.Person{FirstName: "Alex", LastName: "Anderson"}
	// fmt.Println(alex)

	var budi structs.Person
	budi.FirstName = "Budi"
	budi.LastName = "Setiawan"
	// fmt.Printf("%+v", budi)

	jim := structs.Person{
		FirstName: "Jim",
		LastName:  "Office",
		Contact: structs.ContactInfo{
			Email:   "jim@gmail.com",
			ZipCode: 95000,
		},
	}

	fmt.Printf("%+v", jim)

}
