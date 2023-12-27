package main

import "fmt"

func main() {
	type NoKTP string

	var contoh string = "22222222"
	var contohKtp NoKTP = NoKTP("090909090")

	fmt.Println(contoh)
	fmt.Println(contohKtp)
}
