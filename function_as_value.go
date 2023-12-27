package main

import "log"

func getGoodbye(name string) string {
	return "Good bye " + name
}

func main() {
	goodbye := getGoodbye

	log.Println(goodbye("Nael"))
}
