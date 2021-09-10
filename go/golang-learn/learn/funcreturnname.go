package main

import "fmt"

//func getTheName is multiple type data you can do this
// func getTheName() (firstName string, middleName int, lastName string){

// }
func getTheName() (firstName, middleName, lastName string) {
	firstName = "Made"
	middleName = "Garda"
	lastName = "Setiawan"

	return

}

func main() {
	a, b, c := getTheName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
