package main

import (
	"fmt"
	"log"
)

func logging() {
	fmt.Println("Finish")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func endApp() {
	fmt.Println("End app")
	message := recover() // recover should be called inside of a defered function
	// recover will continue the app even if a panic occured
	// similar with javascript's try catch
	log.Println("Error occured:", message)
}

func runApp(error bool) {
	defer endApp() //  defer will always be called even if there are panic
	if error {
		panic("error")
	}
}

func main() {
	runApplication()
	runApp(true)
	log.Println("Unreached code if panic happens without recover")
}