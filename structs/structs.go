package structs

import "fmt"

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
