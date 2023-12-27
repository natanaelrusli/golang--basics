package main

import "fmt"

func main() {
	name := "Nael"

	if name == "Nael" {
		fmt.Println("Hello, Nael")
	} else if name == "Nata" {
		fmt.Println("Hello, Nata")
	} else if name == "Nata" {
		fmt.Println("Hello, Nata lagi")
	} else {
		fmt.Println("Hi!, let's know each other!")
	}

	// if with short statement
	// length := len(name)
	// if length > 3 {
	// 	fmt.Println("4 chars")
	// }
	if length := len(name); length > 3 {
		fmt.Println("4 chars")
	}

	// switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sesuai")
	}

	// switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah sesuai")
	}
}
