package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Made",
		MiddleName: "Garda",
		LastName:   "Setiawan",
		Age: 20,
		Maried: false,
		Hobbies: []string{"gaming","coding","traveling","cooking"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Made","MiddleName":"Garda","LastName":"Setiawan","Age":20,"Maried":false,"Hobbies":["gaming","coding","traveling","cooking"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Made",
		Addresses: []Address{
			{
				Street:     "Jalan Belum ada",
				Country:    "indonesia",
				PostalCode: "08080",
			}, {
				Street:     "Jalan setapak",
				Country:    "usa",
				PostalCode: "3306",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))

}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Made","MiddleName":"","LastName":"","Age":0,"Maried":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum ada","Country":"indonesia","PostalCode":"08080"},{"Street":"Jalan setapak","Country":"usa","PostalCode":"3306"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {

	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
