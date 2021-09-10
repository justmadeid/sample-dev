package main

import (
	"fmt"
	"golang-learn/helper"
)

func main() {
	helper.SayHello("made")
	// result := helper.SayHello("made")
	// fmt.Println(result)

	//tidak bisa di akses karna modifier kecil
	// helper.sayGoodBye("made")
	//fmt.Println(helper.version)

	fmt.Println(helper.Application)
}
