package main

import (
	"errors"
	"log"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	log.Println(Pembagian(1, 0))
	log.Println(Pembagian(10, 2))

	hasil, err := Pembagian(100, 0) 
	if err == nil {
		log.Println(hasil)
	} else {
		log.Println(err)
	}
}
