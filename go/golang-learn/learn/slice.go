package main

import "fmt"

func main() {
	months := [...]string{
		"januari",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubah"
	// fmt.Println(slice1)

	slice1[0] = "gogatsu"
	fmt.Println(months)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "made")
	fmt.Println(slice3)
	slice3[1] = "dessss"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Made"
	newSlice[1] = "Sanskara"

	fmt.Println(newSlice)
	// fmt.Println("lenght dari new slice", len(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//perbedaan array dan slice
	//array ada value didalam bracket/jika tidak mengetahui value nya bisa menggunakan ...
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
