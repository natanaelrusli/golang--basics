package main

import "fmt"

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	names := [...]string{"Nata", "Nael", "Evelyn", "Viriya"}

	slice := names[1:3]
	fmt.Println(slice)
	
	slice2 := names[2:4]
	fmt.Println(slice2)

	slice3 := names[:]
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice2))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Nael"
	newSlice[1] = "Nael"
	// newSlice[2] = "Nael"

	newSlice2 := append(newSlice, "Nael")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice2)

	newSlice2[0] = "Asep"
	fmt.Println(newSlice2)
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	// Copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	// array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// array or slice?
}
