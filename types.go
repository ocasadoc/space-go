package main

import (
	"time"
)

const (
	scheduledLaunches = "s"
)

// UpcomingLaunches matches the JSON returned from the SpaceX API
type UpcomingLaunches []struct {
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
