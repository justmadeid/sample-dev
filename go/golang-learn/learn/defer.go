package main

import "fmt"

func logging() {
	fmt.Println("Success Login")
}

func runApplication(value int) {
	fmt.Println("run Application")
	result := 10 / value
	fmt.Println(result)
	logging()
}

func main() {
	defer logging()
	runApplication(0)
}
