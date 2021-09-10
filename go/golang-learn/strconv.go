package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("flase")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("error", err.Error())
	}
}
