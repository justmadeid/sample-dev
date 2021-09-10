package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("made")

	if person["name"] == "" {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
	// fmt.Println(person)
}
