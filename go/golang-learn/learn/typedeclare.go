package main

import "fmt"

func main() {
	type noKtp string
	type maried bool

	var identitas noKtp = "3849384938493"
	fmt.Println(identitas)

	var heHasMaried maried = true
	fmt.Println(heHasMaried)

}
