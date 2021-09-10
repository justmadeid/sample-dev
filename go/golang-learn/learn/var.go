package main

import "fmt"

func main() {
	var name string

	name = "Made Garda"
	fmt.Println(name)

	name = "Made Sanskara"
	fmt.Println(name)

	var nickname = "Sanskara"
	fmt.Println(nickname)

	var age = 30
	fmt.Println(age)

	// untuk deklarasi selain var menggunakan :=
	country := "Indonesia"
	fmt.Println(country)
	// jika menggunakan var yang sama yang sudah di deklarasikan
	country = "Australia"
	fmt.Println(country)

	// untuk lebih rapih

	var (
		firstname = "Made"
		lastname  = "sanskara"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
