package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Maried() {
	man.Name = "Mr . " + man.Name
}

func main() {
	user := Man{"Made"}
	user.Maried()

	fmt.Println(user.Name)
}
