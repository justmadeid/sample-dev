package main

import (
	"fmt"
	"golang-learn/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
