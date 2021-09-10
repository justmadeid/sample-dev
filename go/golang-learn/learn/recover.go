package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("terjadi error dengan pesan", message)
	}
	fmt.Println("end application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("0X22271")
	}
	fmt.Println("app berjalan")
}

func main() {
	runApp(true)
}
