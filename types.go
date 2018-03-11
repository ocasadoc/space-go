package main

const (
	upcomingLaunches = "s"
	spaceAgencies    = "a"
)

// UpcomingLaunches matches the JSON returned by the launchlibrari API
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

// SpaceAgencies matches the JSON returned by the launchlibrari API
type SpaceAgencies []struct {
	Agencies []struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		CountryCode string        `json:"countryCode"`
		Abbrev      string        `json:"abbrev"`
		Type        int           `json:"type"`
		InfoURL     string        `json:"infoURL"`
		WikiURL     string        `json:"wikiURL"`
		InfoURLs    []interface{} `json:"infoURLs"`
		Islsp       int           `json:"islsp"`
	} `json:"agencies"`
	Total  int `json:"total"`
	Count  int `json:"count"`
	Offset int `json:"offset"`
}
