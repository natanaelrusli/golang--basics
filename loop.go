package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Iteration", counter)
	// 	counter++
	// }

	
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Iteration", counter)
	}

	fmt.Println("Finished")

	names := []string{"Nata", "Nael"}

	for i, val := range names {
		fmt.Println(i, val)
	}

	// break and continue
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}

		if i == 9 {
			break
		}

		fmt.Println("Iteration", i)
	}
}
