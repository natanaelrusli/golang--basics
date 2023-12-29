package main

import "log"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "Nael" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	log.Println(SaveData("", nil))

	// checking error types
	err := SaveData("1", nil)
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			log.Println("validation error:", validationError.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			log.Println("not found error:", notFoundError.Message)
		} else {
			log.Println("unknown error:", err.Error())
		}

		// theese conditional can be refactored to switch cases
	}
}