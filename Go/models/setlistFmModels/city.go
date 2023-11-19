package setlistFmModels

type City struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	StateCode string  `json:"stateCode"`
	State     string  `json:"state"`
	Coords    Coords  `json:"coords"`
	Country   Country `json:"country"`
}
