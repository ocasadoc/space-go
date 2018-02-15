package main

import (
	"time"
)

const (
	scheduledLaunches = "s"
)

// UpcomingLaunches matches the JSON returned from the SpaceX API
type UpcomingLaunches struct {
	Launches []struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		Windowstart string        `json:"windowstart"`
		Windowend   string        `json:"windowend"`
		Net         string        `json:"net"`
		Wsstamp     int           `json:"wsstamp"`
		Westamp     int           `json:"westamp"`
		Netstamp    int           `json:"netstamp"`
		Isostart    string        `json:"isostart"`
		Isoend      string        `json:"isoend"`
		Isonet      string        `json:"isonet"`
		Status      int           `json:"status"`
		Inhold      int           `json:"inhold"`
		Tbdtime     int           `json:"tbdtime"`
		VidURLs     []string      `json:"vidURLs"`
		VidURL      interface{}   `json:"vidURL"`
		InfoURLs    []interface{} `json:"infoURLs"`
		InfoURL     interface{}   `json:"infoURL"`
		Holdreason  interface{}   `json:"holdreason"`
		Failreason  interface{}   `json:"failreason"`
		Tbddate     int           `json:"tbddate"`
		Probability int           `json:"probability"`
		Hashtag     interface{}   `json:"hashtag"`
		Location    struct {
			Pads []struct {
				ID        int     `json:"id"`
				Name      string  `json:"name"`
				InfoURL   string  `json:"infoURL"`
				WikiURL   string  `json:"wikiURL"`
				MapURL    string  `json:"mapURL"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Agencies  []struct {
					ID          int         `json:"id"`
					Name        string      `json:"name"`
					Abbrev      string      `json:"abbrev"`
					CountryCode string      `json:"countryCode"`
					Type        int         `json:"type"`
					InfoURL     interface{} `json:"infoURL"`
					WikiURL     string      `json:"wikiURL"`
					InfoURLs    []string    `json:"infoURLs"`
				} `json:"agencies"`
			} `json:"pads"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			InfoURL     string `json:"infoURL"`
			WikiURL     string `json:"wikiURL"`
			CountryCode string `json:"countryCode"`
		} `json:"location"`
		Rocket struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Configuration string `json:"configuration"`
			Familyname    string `json:"familyname"`
			Agencies      []struct {
				ID          int         `json:"id"`
				Name        string      `json:"name"`
				Abbrev      string      `json:"abbrev"`
				CountryCode string      `json:"countryCode"`
				Type        int         `json:"type"`
				InfoURL     interface{} `json:"infoURL"`
				WikiURL     string      `json:"wikiURL"`
				InfoURLs    []string    `json:"infoURLs"`
			} `json:"agencies"`
			WikiURL    string   `json:"wikiURL"`
			InfoURLs   []string `json:"infoURLs"`
			InfoURL    string   `json:"infoURL"`
			ImageSizes []int    `json:"imageSizes"`
			ImageURL   string   `json:"imageURL"`
		} `json:"rocket"`
		Missions []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        int    `json:"type"`
			TypeName    string `json:"typeName"`
			Agencies    []struct {
				ID          int         `json:"id"`
				Name        string      `json:"name"`
				Abbrev      string      `json:"abbrev"`
				CountryCode string      `json:"countryCode"`
				Type        int         `json:"type"`
				InfoURL     interface{} `json:"infoURL"`
				WikiURL     string      `json:"wikiURL"`
				InfoURLs    []string    `json:"infoURLs"`
			} `json:"agencies"`
		} `json:"missions"`
		Lsp struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Abbrev      string      `json:"abbrev"`
			CountryCode string      `json:"countryCode"`
			Type        int         `json:"type"`
			InfoURL     interface{} `json:"infoURL"`
			WikiURL     string      `json:"wikiURL"`
			InfoURLs    []string    `json:"infoURLs"`
		} `json:"lsp"`
	} `json:"launches"`
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

// UpcomingLaunchesSpaceX matches the JSON returned from the SpaceX API
type UpcomingLaunchesSpaceX []struct {
	FlightNumber    int       `json:"flight_number"`
	LaunchYear      string    `json:"launch_year"`
	LaunchDateUnix  int       `json:"launch_date_unix"`
	LaunchDateUtc   time.Time `json:"launch_date_utc"`
	LaunchDateLocal string    `json:"launch_date_local"`
	Rocket          struct {
		RocketID   string `json:"rocket_id"`
		RocketName string `json:"rocket_name"`
		RocketType string `json:"rocket_type"`
		FirstStage struct {
			Cores []struct {
				CoreSerial     string      `json:"core_serial"`
				Reused         bool        `json:"reused"`
				LandSuccess    bool        `json:"land_success"`
				LandingType    interface{} `json:"landing_type"`
				LandingVehicle interface{} `json:"landing_vehicle"`
			} `json:"cores"`
		} `json:"first_stage"`
		SecondStage struct {
			Payloads []struct {
				PayloadID      string      `json:"payload_id"`
				Reused         bool        `json:"reused"`
				Customers      []string    `json:"customers"`
				PayloadType    string      `json:"payload_type"`
				PayloadMassKg  int         `json:"payload_mass_kg"`
				PayloadMassLbs interface{} `json:"payload_mass_lbs"`
				Orbit          string      `json:"orbit"`
			} `json:"payloads"`
		} `json:"second_stage"`
	} `json:"rocket"`
	Telemetry struct {
		FlightClub interface{} `json:"flight_club"`
	} `json:"telemetry"`
	Reuse struct {
		Core      bool `json:"core"`
		SideCore1 bool `json:"side_core1"`
		SideCore2 bool `json:"side_core2"`
		Fairings  bool `json:"fairings"`
		Capsule   bool `json:"capsule"`
	} `json:"reuse"`
	LaunchSite struct {
		SiteID       string `json:"site_id"`
		SiteName     string `json:"site_name"`
		SiteNameLong string `json:"site_name_long"`
	} `json:"launch_site"`
	LaunchSuccess interface{} `json:"launch_success"`
	Links         struct {
		MissionPatch   interface{} `json:"mission_patch"`
		RedditCampaign string      `json:"reddit_campaign"`
		RedditLaunch   interface{} `json:"reddit_launch"`
		RedditRecovery interface{} `json:"reddit_recovery"`
		RedditMedia    interface{} `json:"reddit_media"`
		Presskit       interface{} `json:"presskit"`
		ArticleLink    interface{} `json:"article_link"`
		VideoLink      interface{} `json:"video_link"`
	} `json:"links"`
	Details interface{} `json:"details"`
}
