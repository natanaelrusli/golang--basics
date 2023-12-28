package main

import "log"

func random() interface{} {
	return "OK"
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	var result any = random()
	resultString, ok := result.(string)
	if ok {
		log.Println(resultString)
	}

	resultInt, ok := result.(int) // panic
	if ok {
		log.Println(resultInt)
	} else {
		log.Println("Conversion failed")
	}

	// assertions switch
	switch value := result.(type) {
	case string:
		log.Println("String", value)
	case int:
		log.Println("Int", value)
	default:
		log.Println("Unknown type")
	}
}