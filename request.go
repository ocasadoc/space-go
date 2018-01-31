package main

import (
	"log"
	"net/http"
)

func getUpcomingLaunches() {
	url := "https://api.spacexdata.com/v2/launches/upcoming"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request creation failed: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP request failed: ", err)
		return
	}

	defer resp.Body.Close()

	drawTable(resp.Body, scheduledLaunches)
	askForLaunch()
}
