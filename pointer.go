package main

import "log"

// in go, when we send a variable to a function, method or other variable, what actually being send is the duplicate of the value

type Address struct {
	City string
	Province string
	Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	log.Println(address1)
	log.Println(*address2)
}
