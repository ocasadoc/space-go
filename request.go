package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getUpcomingLaunches() (UpcomingLaunches, error) {
	url := "https://api.spacexdata.com/v2/launches/upcoming"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request creation failed: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP request failed: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		fmt.Println("Error decoding json")
	}

	return record, nil
}
