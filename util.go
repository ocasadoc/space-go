package main

import (
	"fmt"
)

func getUserInput() string {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n")
	return response
}
