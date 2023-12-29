package main

import "log"

type Address struct {
	City string
	Province string
	Country string
}

// ** Without Pointer **
// func ChangeAddressToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

// func main() {
// 	var address Address = Address{"Subang", "Jawa Barat", ""}
// 	ChangeAddressToIndonesia(address)

// 	log.Println(address)
// }

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// var address *Address = &Address{"Subang", "Jawa Barat", ""}
	// ChangeAddressToIndonesia(address)
	
	var address Address = Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	log.Println(address)
}