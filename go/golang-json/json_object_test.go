package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Maried bool
	Hobbies []string
	Addresses []Address
}


func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Made",
		MiddleName: "Garda",
		LastName:   "Setiawan",
		Age: 20,
		Maried: false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}