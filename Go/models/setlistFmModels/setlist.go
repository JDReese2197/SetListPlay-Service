package setlistFmModels

type Setlist struct {
	Artist      Artist `json:"artist"`
	Venue       Venue  `json:"venue"`
	Tour        Tour   `json:"tour"`
	Set         []Set  `json:"set"`
	Info        string `json:"info"`
	Url         string `json:"url"`
	Id          string `json:"id"`
	VersionId   string `json:"versionId"`
	EventDate   string `json:"eventDate"`
	LastUpdated string `json:"lastUpdated"`
}
