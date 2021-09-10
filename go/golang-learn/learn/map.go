package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Made",
		"address": "Semarang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "Made Garda"
	book["edition"] = "revision"

	fmt.Println(book, len(book))
	delete(book, "edition")
	fmt.Println(book, len(book))

}
