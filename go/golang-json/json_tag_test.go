package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P001",
		Name:     "Apple Mac",
		ImageURL: "http://jsrmade.me",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"Apple Mac","image_url":"http://jsrmade.me"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes,product)
	fmt.Println(product)

}