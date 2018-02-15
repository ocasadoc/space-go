package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	printWelcomeMessage()

	scheduleArg := os.Args[1]
	if scheduleArg == "s" {
		printUpcomingLaunches()
	}
}

func printWelcomeMessage() {
	fmt.Println("Welcome to space-go!")
	fmt.Println("" +
		"         __\n" +
		"         \\ \\_____\n" +
		"*********[==_____>\n" +
		"         /_/\n")
}

func printUpcomingLaunches() {
	fmt.Println("The next scheduled launches are:")
	upcomingLaunches, err := getUpcomingLaunches()
	if err != nil {
		log.Fatal("Request creation failed: ", err)
		return
	}
	printRetrievedLaunches(upcomingLaunches)
	askForLaunchDetails()
}
