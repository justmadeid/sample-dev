package main

import "fmt"

func main() {
	name := "Made"

	if name == "Made" {
		fmt.Println("Hello Admin", name)
	} else if name == "Sanskara" {
		fmt.Println("Hello User", name)
	} else {
		fmt.Println("Nama ", name, "Tidak Terdaftar")
	}

	//short statement if
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("pass")
	}

	//explore
	people := [...]string{
		"budi",
		"a",
	}

	if len(people[1]) > 2 {
		slice := people[:]
		fmt.Println(slice, " Sebelum di ubah")
		sliceApp := append(slice, "ok")
		sliceApp[0] = "si"
		fmt.Println(sliceApp, "setelah Di ubah")

		copySlice := make([]string, len(sliceApp), cap(sliceApp))
		copy(copySlice, sliceApp)
		fmt.Println("hasil copy", copySlice)

	} else {
		newMap := make(map[string]string)
		newMap["nama"] = "made"
		newMap["status"] = "Pelajar"

		fmt.Println(newMap)

		delete(newMap, "status")

		fmt.Println(newMap)

	}

}
