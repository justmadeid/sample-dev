package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "made"
	names[1] = "garda"
	names[2] = "setiawan"

	fmt.Println(names[0], names[1])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println("===================")

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}
