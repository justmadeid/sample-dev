package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// var addsress1 Address = Address{"Subang", "Jawa Barat", "indonesia"}
	addsress1 := Address{"Subang", "Jawa Barat", "indonesia"}
	addsress2 := &addsress1
	//atau dngan pendekalasian seperti ini
	var addsress3 *Address = &addsress1

	addsress2.City = "Bandung"

	*addsress2 = Address{"malang", "jawa timur", "indonesia"}

	fmt.Println(addsress1)
	fmt.Println(addsress2)
	fmt.Println(addsress3)

	var addsress4 *Address = new(Address)
	// addsress4 := new(Address)
	addsress4.City = "semarang" //kalo gak di isi menjadi kosong
	fmt.Println(addsress4)

	var alamat = Address{
		City:     "Subang",
		Province: "jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	// ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
