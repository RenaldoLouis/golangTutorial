package structs

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   ContactInfo
}

type ContactInfo struct {
	Email   string `json:"Email"`
	ZipCode int64  `json:"Zip"`
}
