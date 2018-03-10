package main

import (
	"fmt"
	"github.com/apcera/termtables"
	"github.com/fatih/color"
	"log"
	"strconv"
)

var numberOfLaunches int
var record UpcomingLaunches

// TableTypes represents the different tables that can be printed in our program
type TableTypes string

func printFetchedLaunches(upcomingLaunches UpcomingLaunches) {
	drawTable(upcomingLaunches, scheduledLaunches)
}

func drawTable(upcomingLaunches UpcomingLaunches, tableType TableTypes) {
	record = upcomingLaunches
	switch tableType {
	case scheduledLaunches:
		drawScheduledLaunchesTable()
	default:
		log.Fatal("Error while drawing table")
	}
}

func drawScheduledLaunchesTable() {
	table := termtables.CreateTable()
	table.AddHeaders("N", "Name", "Net", "tbdtime", "tbddate")
	color.Set(color.FgYellow)
	for i := 0; i < len(record.Launches); i++ {
		table.AddRow(i, record.Launches[i].Name, record.Launches[i].Net,
			record.Launches[i].Tbdtime, record.Launches[i].Tbddate)
	}
	numberOfLaunches = len(record.Launches)
	color.Unset()

	fmt.Println(table.Render())
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
	prettyPrintJSON(record.Launches[launchNumber])
}
