package main

import "fmt"

func main() {
	person := map[string]string {
		"name": "Nael", 
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["kosong"])

	// making a new map
	book := make(map[string]string)
	book["title"] = "Buku golang"
	book["author"] = "Nael"
	book["wrong"] = "Salah"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
