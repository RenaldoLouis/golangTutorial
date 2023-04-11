package structs

import "fmt"

type bot interface {
	GetGreeting() string
}

type EnglishBot struct{}

type SpanishBot struct{}

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	ContactInfo
}

type ContactInfo struct {
	Email   string `json:"Email"`
	ZipCode int64  `json:"Zip"`
}

func (p Person) PrintString() {
	fmt.Printf("%+v", p)
}

func (p *Person) UpdateName(newFirstName string) {
	p.FirstName = newFirstName
}

func (eb EnglishBot) GetGreeting() string {
	return "Hi There"
}

func (sb SpanishBot) GetGreeting() string {
	return "Hola"
}

func PrintGreeting(b bot) {
	fmt.Println(b.GetGreeting())
}
