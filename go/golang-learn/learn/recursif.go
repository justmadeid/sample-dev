package main

import "fmt"

//using for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//recursive func
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}

}

func main() {
	//result for loop
	loop := factorialLoop(5)
	fmt.Println("lopp", loop)
	//result racursive
	recursive := factorialRecursive(5)
	fmt.Println("recursive", recursive)
}
