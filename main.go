package main

import (
	"fmt"
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
		handleUserSelection(os.Args[1])
	} else {
		printMainMenu()
	}
}

func printMainMenu() {
	fmt.Printf("Type the option you prefer:"+
		"\n%s: Upcoming launches\n",
		upcomingLaunches)
	fmt.Printf(" > ")
	handleUserSelection(getUserInput())
}

func handleUserSelection(option string) {
	switch option {
	case upcomingLaunches:
		printUpcomingLaunches()
	default:
		printMainMenu()
	}
}

func printUpcomingLaunches() {
	fmt.Println("The upcoming launches are:")
	upcomingLaunches, err := fetchUpcomingLaunches()
	if err != nil {
		fmt.Println(err)
		return
	}
	printFetchedLaunches(upcomingLaunches)
	askForLaunchDetails()
}
