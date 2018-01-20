package main

import (
	"encoding/json"
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

	var record UpcomingLaunches

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	println("flightNumber: ", record[0].FlightNumber)
	println("launchDate: ", record[0].LaunchDateUtc.String())
	println("launchDate: ", record[0].Rocket.RocketName)

}
