package main

import "fmt"

type BalckList func(string) bool

func registerUser(name string, balckList BalckList) {
	if balckList(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//func dimasukan ke dalam variable
	balckList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", balckList)
	registerUser("made", balckList)

	//func langsung kedalam variable
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
