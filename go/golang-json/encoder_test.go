package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Made",
		MiddleName: "Garda",
		LastName:   "Setiawan",
		Age: 20,
		Maried: false,
	}

	_ = encoder.Encode(customer)
	fmt.Println(customer)

}
