package main

import (
	"os"
)

func main() {

	scheduleArg := os.Args[1]

	if scheduleArg == "s" {
		println("Welcome to space-go!")
		println("The next scheduled launch is:")
		getUpcomingLaunches()
	}
}
