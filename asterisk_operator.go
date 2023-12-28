package main

import "log"


type Address struct {
	City string
	Province string
	Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	log.Println(address1)
	log.Println(*address2)

	// memindahkan semua variable yg reference ke Address sebelumnya ke Address yang baru
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	log.Println(address1)
	log.Println(address2)
	log.Println(address3)
}