package main

import (
	"belajar-golang-dasar/database"
	"log"
)

func main() {
	log.Println("DB Connection:", database.GetDatabaseConnection())
}