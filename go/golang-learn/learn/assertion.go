package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string ", value)
	case int:
		fmt.Println("int ", value)
	default:
		fmt.Println("unkown")
	}
}
