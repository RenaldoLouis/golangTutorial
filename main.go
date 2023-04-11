package main

import (
	"fmt"
	decks "golangTutorial/deck"
	"golangTutorial/structs"
	"net/http"
	"time"
)

var (
	decksVar decks.InterfaceDecks = decks.NewDecks()
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("len", len(bs))

	return len(bs), nil
}

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

	// jim := structs.Person{
	// 	FirstName: "Jim",
	// 	LastName:  "Office",
	// 	ContactInfo: structs.ContactInfo{
	// 		Email:   "jim@gmail.com",
	// 		ZipCode: 95000,
	// 	},
	// }

	// jim.UpdateName("Tommy")
	// jim.PrintString()

	//Part 3 Maps
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#4bf745",
	// 	"white": "#fffff",
	// }

	// printMap(colors)

	//Part 4 Interfaces

	// eb := structs.EnglishBot{}
	// sb := structs.SpanishBot{}

	// structs.PrintGreeting(eb)
	// structs.PrintGreeting(sb)
	// resp, err := http.Get("http://google.com")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	//Part 5 Channels
	links := []string{
		"http://amazon.com",
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://twitter.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
