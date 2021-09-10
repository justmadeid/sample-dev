package main

import "fmt"

func sumAll(numbers ...int) int {

	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 90, 90, 90)
	fmt.Println(total)

	//memasukan slice ke func variadic
	slice := []int{10, 10, 10, 10, 10}
	//jika slice harus menambahkan ...
	total = sumAll(slice...)
	fmt.Println(total)

}
