package main

import "fmt"

func main() {
	var i = 10

	// augmented assignment
	i += 10
	fmt.Println(i)

	i += 20
	fmt.Println(i)

	// unary operator
	i++
	i--
	i--
	fmt.Println(i)
}