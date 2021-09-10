package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {

	return len(value)

}

func (value UserSlice) Less(i, j int) bool {

	return value[i].Age < value[j].Age

}

func (value UserSlice) Swap(i, j int) {

	value[i], value[j] = value[j], value[i]

}

func main() {
	users := []User{
		{"made", 20},
		{"m", 21},
		{"a", 22},
		{"d", 23},
		{"e", 24},
	}
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
