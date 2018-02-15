package main

import (
	"encoding/json"
	"fmt"
	"github.com/apcera/termtables"
	"github.com/fatih/color"
	"io"
	"log"
	"strconv"
	"strings"
)

var numberOfLaunches int
var record UpcomingLaunches

// TableTypes represents the different tables that can be printed in our program
type TableTypes string

func printRetrievedLaunches(upcomingLaunches UpcomingLaunches) {
	drawTable(upcomingLaunches, scheduledLaunches)
}

func drawTable(upcomingLaunches UpcomingLaunches, tableType TableTypes) {
	record = upcomingLaunches
	switch tableType {
	case scheduledLaunches:
		table := termtables.CreateTable()
		table.AddHeaders("N", "Name", "Net", "tbdtime", "tbddate")
		color.Set(color.FgYellow)
		for i := 0; i < len(record.Launches); i++ {
			// TODO: Hardcoded company name
			table.AddRow(i, record.Launches[i].Name, record.Launches[i].Net,
				record.Launches[i].Tbdtime, record.Launches[i].Tbddate)
		}
		numberOfLaunches = len(record.Launches)
		color.Unset()

		fmt.Println(table.Render())

	default:
		log.Fatal("Error while drawing table")
	}
}

func askForLaunchDetails() {
	var response string
	fmt.Print(" > Type the number of the launch to see more information: ")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	responseConverted, err := strconv.Atoi(response)

	if err != nil {
		if response == "b" {
			printUpcomingLaunches()
		} else {
			fmt.Printf(" >>> Type the number or 'b' to go back\n")
			askForLaunchDetails()
		}
	} else {
		if responseConverted < numberOfLaunches && responseConverted >= 0 {
			getDetailedLaunchInfo(responseConverted)
			askForLaunchDetails()
		} else {
			fmt.Printf(" >>> Type the number or 'b' to go back\n")
			askForLaunchDetails()
		}
	}
}

func getDetailedLaunchInfo(launchNumber int) {
	jsonString, err := json.Marshal(record.Launches[launchNumber])
	if err != nil {
		fmt.Println("Error converting json to")
	}
	decodedJSON := json.NewDecoder(strings.NewReader(string(jsonString)))
	jsonPairFlag, spaces, lineBreakDone := 0, 0, false

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
			if string(delim) == "{" || string(delim) == "[" {
				spaces++
				fmt.Printf("\n")
				lineBreakDone = false
				jsonPairFlag = 0

			}
			if string(delim) == "}" || string(delim) == "]" {
				spaces--
				if !lineBreakDone {
					fmt.Printf("\n")
					lineBreakDone = true
				}
				jsonPairFlag = 0
			}
		} else {
			if jsonPairFlag == 0 {
				for i := 0; i < spaces; i++ {
					fmt.Printf("\t")
				}
				if _, ok := token.(string); ok {
					fmt.Printf("%v: ", strings.Title(token.(string)))
				} else {
					fmt.Printf("%v: ", token.(float64))
				}

				jsonPairFlag++
			} else {
				fmt.Printf("%v\n", token)
				jsonPairFlag--
				lineBreakDone = true
			}
		}
	}
}
