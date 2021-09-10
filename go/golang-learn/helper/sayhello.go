package helper

import "fmt"

var version = "1.0.0.1"
var Application = "Golang"

func SayHello(name string) {
	fmt.Println("Hello ", name)
}

func sayGoodBye(name string) {
	fmt.Println("bye ", name)
}

// func SayHello(name string)string  {
// 	return "Hallo " + name
// }
