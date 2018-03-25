package main

import (
	"encoding/json"
	"fmt"
)

func prettyPrintJSON(v interface{}) {
	json, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(json))
}
