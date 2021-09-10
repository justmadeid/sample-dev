package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Made", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
