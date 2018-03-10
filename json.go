package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func prettyPrintJSON(v interface{}) {
	jsonString, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error converting json to")
	}
	decodedJSON := json.NewDecoder(strings.NewReader(string(jsonString)))

	// loop until EOF
	for {
		token, err := decodedJSON.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		delim, ok := token.(json.Delim)
		if err != nil {
			fmt.Println("Error converting token to string")
		}
		if ok {
			handleJSONDelimiter(string(delim))
		} else {
			handleJSONToken(token)
		}
	}
}

var jsonPairFlag = 0
var spaces = 0
var lineBreakDone = false

func handleJSONDelimiter(delim string) {
	if delim == "{" || delim == "[" {
		spaces++
		lineBreakDone = true
	}
	if delim == "}" || delim == "]" {
		spaces--
		if lineBreakDone {
			fmt.Printf("\n")
		}
		lineBreakDone = false
	}
}

func handleJSONToken(token interface{}) {
	if lineBreakDone {
		for i := 0; i < spaces; i++ {
			fmt.Printf("\t")
		}
		if _, ok := token.(string); ok {
			fmt.Printf("%v: ", token.(string))
		} else {
			fmt.Printf("%v: ", token.(float64))
		}
	} else {
		fmt.Printf("%v", token)
	}
}
