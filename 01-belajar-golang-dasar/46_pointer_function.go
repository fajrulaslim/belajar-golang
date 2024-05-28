package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {	
	// cara1
	// var address *Address = &Address{}
	// ChangeCountryToIndonesia(address)

	// cara2
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
