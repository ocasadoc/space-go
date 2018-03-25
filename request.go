package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchUpcomingLaunches() (UpcomingLaunches, error) {
	url := "https://launchlibrary.net/1.3/launch?mode=verbose"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return UpcomingLaunches{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return UpcomingLaunches{}, err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		fmt.Println(err)
	}

	return record, nil
}
