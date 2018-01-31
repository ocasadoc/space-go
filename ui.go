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

func drawTable(body io.ReadCloser, tableType TableTypes) {
	switch tableType {
	case scheduledLaunches:
		table := termtables.CreateTable()
		table.AddHeaders("N", "Company", "LaunchDate", "Rocket name", "Rocket type", "Launch success", "Core reuse", "Launch site")
		if err := json.NewDecoder(body).Decode(&record); err != nil {
			fmt.Println("Error decoding json")
		}
		color.Set(color.FgYellow)
		for i := 0; i < len(record); i++ {
			// TODO: Hardcoded company name
			table.AddRow(i, "SpaceX", record[i].LaunchDateUtc, record[i].Rocket.RocketName,
				record[i].Rocket.RocketType, record[i].LaunchSuccess,
				record[i].Reuse.Core, record[i].LaunchSite.SiteName)
		}
		numberOfLaunches = len(record)
		color.Unset()

		fmt.Println(table.Render())

	default:
		log.Fatal("Error while drawing table")
	}
}

func askForLaunch() {
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
			askForLaunch()
		}
	} else {
		if responseConverted < numberOfLaunches && responseConverted >= 0 {
			getDetailedLaunchInfo(responseConverted)
			askForLaunch()
		} else {
			fmt.Printf(" >>> Type the number or 'b' to go back\n")
			askForLaunch()
		}
	}
}

func getDetailedLaunchInfo(launchNumber int) {
	jsonString, err := json.Marshal(record[launchNumber])
	if err != nil {
		fmt.Println("Error converting json to")
	}
	decodedJson := json.NewDecoder(strings.NewReader(string(jsonString)))
	jsonPairFlag, spaces, lineBreakDone := 0, 0, false

	for {
		token, err := decodedJson.Token()
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
				fmt.Printf("%v: ", strings.Title(strings.Replace(token.(string), "_", " ", -1)))
				jsonPairFlag++
			} else {
				fmt.Printf("%v\n", token)
				jsonPairFlag--
				lineBreakDone = true
			}
		}
	}
}
