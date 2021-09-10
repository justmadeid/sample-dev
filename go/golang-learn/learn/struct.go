package main

import "fmt"

type Customer struct {
	Name, Address string
	Balance       int
	Active        bool
}

//struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hai", name, "my name is", customer.Name)
}

func (a Customer) huuu() {
	fmt.Println("huuuu", a.Name)
}

func main() {
	var user Customer
	user.Name = "Made Garda Setiawan"
	user.Address = "Semarang"
	user.Balance = 100000

	//memanggil struct method
	// sayHi(user, "sanskara")
	user.sayHi("Sanskara")
	user.huuu()

	fmt.Println(user.Name)
	fmt.Println(user.Address)
	fmt.Println(user.Balance)
	fmt.Println(user)

	// struct literal

	//di sarankan
	store := Customer{
		Name:    "Makanku",
		Address: "Semarang",
		Balance: 1000000,
	}

	fmt.Println(store)

	//tidak disarankan akan banyak masalah di kedepannya
	market := Customer{"Minumku", "Tembalang", 500000, true}

	fmt.Println(market)

}
