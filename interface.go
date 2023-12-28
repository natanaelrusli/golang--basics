package main

import "log"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(hasName HasName) {
	log.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups() interface{} {
	// return 1
	// return true
	return "Ups"
}

func main() {
	person := Person{Name: "Nael"}
	animal := Animal{Name: "Hoshi"}
	
	SayHello(person)
	SayHello(animal)

	var kosong any = Ups()
	log.Println(kosong)
}