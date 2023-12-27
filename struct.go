package main

import "log"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var nael Customer
	nael.Name = "Nata Nael"
	nael.Age = 25
	nael.Address = "Jakarta"

	log.Println(nael)
	log.Println(nael.Name)
	log.Println(nael.Address)
	log.Println(nael.Age)

	// evelyn := Customer{
	// 	Name: "Evelyn Viriya",
	// 	Address: "Kalideres",
	// 	Age: 22,
	// }

	evelyn := Customer{"Evelyn Viriya", "Kalideres", 22}

	log.Println(evelyn)
}
