package main

import "fmt"

func main() {
	name := "made"
	counter := 0
	increment := func() {
		name = "oi"
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
