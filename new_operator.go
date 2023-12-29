package main

import "log"

type Address struct {
	City string
	Province string
	Country string
}

func main() {
	// alamat1 := &Address{}
	alamat1 := new(Address) // return a pointer to an empty Address
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	log.Println(alamat1)
	log.Println(alamat2)
}