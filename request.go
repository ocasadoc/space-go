package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func fetchUpcomingLaunches() (UpcomingLaunches, error) {
	url := "https://launchlibrary.net/1.3/launch?mode=verbose"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Request creation failed: ", err)
		return UpcomingLaunches{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("HTTP request failed: ", err)
		return UpcomingLaunches{}, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatal("Error decoding json")
	}

	return record, nil
}
