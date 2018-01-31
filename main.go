package main

import (
	"fmt"
	"os"
)

func main() {

	scheduleArg := os.Args[1]

	if scheduleArg == "s" {
		fmt.Println("Welcome to space-go!")
		fmt.Println("" +
			"   __\n" +
			"   \\ \\_____\n" +
			"***[==_____>\n" +
			"   /_/\n")
		printUpcomingLaunches()
	}
}

func printUpcomingLaunches() {
	fmt.Println("The next scheduled launches are:")
	getUpcomingLaunches()
}
