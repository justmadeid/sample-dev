package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	//for init statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("perulangan ke-", counter2)
	}

	slice := []string{
		"made",
		"sanskara",
		"agastya",
	}

	// fmt.Println(slice[0])
	// fmt.Println(slice[1])
	// fmt.Println(slice[2])

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "made"
	person["title"] = "programmer"

	for key, v := range person {
		fmt.Println(key, "=", v)
	}

}
