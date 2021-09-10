package main

import "fmt"

func getFullName() (string, string, string) {
	return "Made", "Sanskara", "Cool"
}

func main() {
	firstName, _, tagline := getFullName()
	fmt.Println(firstName, tagline)
}
