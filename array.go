package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Nata"
	names[1] = "Nael"

	fmt.Println(names)

	var values = [...]int{
		1,
		2,
	}

	fmt.Println(values)

	fmt.Println("len:", len(values))
}
