package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	printWelcomeMessage()
	handleProgramArgs()
}

func printWelcomeMessage() {
	fmt.Println("Welcome to space-go!")
	fmt.Println("" +
		"         __\n" +
		"         \\ \\_____\n" +
		"*********[==_____>\n" +
		"         /_/\n")
}

func handleProgramArgs() {
	if len(os.Args) == 2 {
		switch option := os.Args[1]; option {
		case "s":
			printUpcomingLaunches()
		case "a":
			//printSpaceAgencies()
		default:
			printHelp()
		}
	} else {
		printHelp()
	}
}

func printUpcomingLaunches() {
	fmt.Println("The next scheduled launches are:")
	upcomingLaunches, err := fetchUpcomingLaunches()
	if err != nil {
		log.Fatal("Request creation failed: ", err)
		return
	}
	printFetchedLaunches(upcomingLaunches)
	askForLaunchDetails()
}

func printHelp() {
	fmt.Printf("Usage:\ns: Upcoming launches\na: Space agencies")
}
