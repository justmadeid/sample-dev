package main

import "fmt"

func main() {

	name := "sanskara"

	switch name {
	case "made":
		fmt.Println("hello made")
	case "sanskara":
		fmt.Println("hello sanskara")
	default:
		fmt.Println("kenalan Dong :>")
	}

	//short statments

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("nama yang luar biasa")
	case false:
		fmt.Println("Nama yang Simple")
	}

	//simple
	leng := len(name)
	switch {
	case len(name) > 10:
		fmt.Println("wow")
	case leng > 5:
		fmt.Println("hey")
	default:
		fmt.Println(":<")
	}

}
