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

func printFetchedLaunches(fetchedUpcomingLaunches UpcomingLaunches) {
	drawTable(fetchedUpcomingLaunches, upcomingLaunches)
}

func drawTable(fetchedUpcomingLaunches UpcomingLaunches, tableType TableTypes) {
	record = fetchedUpcomingLaunches
	switch tableType {
	case upcomingLaunches:
		drawUpcomingLaunchesTable()
	default:
		log.Fatal("Error while drawing table")
	}
}

func drawUpcomingLaunchesTable() {
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
	fmt.Print(" > Type the number of the launch to see more information (or b to go back): ")
	responseString := getUserInput()
	responseInteger, err := strconv.Atoi(responseString)

	if err != nil {
		if responseString == "b" {
			printMainMenu()
		} else {
			askForLaunchDetails()
		}
	} else {
		if responseInteger < numberOfLaunches && responseInteger >= 0 {
			getDetailedLaunchInfo(responseInteger)
			askForLaunchDetails()
		} else {
			askForLaunchDetails()
		}
	}
}

func getDetailedLaunchInfo(launchNumber int) {
	prettyPrintJSON(record.Launches[launchNumber])
}
