package main

import "log"

type Customer struct {
	Name, Address string
	Age int
}

// struct can be used as a parameter of a function
// function inside of a struct called a method

func (customer Customer) sayHello() {
	log.Println("Hello, My name is", customer.Name)
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

	alin := Customer{Name: "Alin"}
	alin.sayHello()
}
