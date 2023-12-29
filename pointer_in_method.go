package main

import "log"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	nael := Man{"Nael"}
	nael.Married()

	log.Println(nael)
}

// for method, it is very recommended to use pointer