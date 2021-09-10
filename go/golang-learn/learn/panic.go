package main

import "fmt"

func endApp() {
	fmt.Println("end application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("app error")
	}
	fmt.Println("app running")
}

func main() {
	runApp(true)
}
