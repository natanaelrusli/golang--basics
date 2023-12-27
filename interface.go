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

func main() {
	person := Person{Name: "Nael"}
	animal := Animal{Name: "Hoshi"}
	
	SayHello(person)
	SayHello(animal)
}