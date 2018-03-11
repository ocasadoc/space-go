package main

import (
	"fmt"
	"log"
)

func getUserInput() string {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")
	return response
}
