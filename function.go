package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Nata"
	middleName = "Nael"
	lastName = "Rusli"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}