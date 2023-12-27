package main

import "log"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	log.Println(sumAll(10, 100, 4, 6, 7, 8, 199))
	log.Println(sumAll(10, 100, 4, 2, 7, 8, 199, 2, 3, 5))

	numbers := []int{10, 10, 10, 10}
	log.Println(sumAll(numbers...))
}