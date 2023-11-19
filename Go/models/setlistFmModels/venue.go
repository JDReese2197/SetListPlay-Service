package setlistFmModels

type Venue struct {
	City City   `json:"city"`
	Url  string `json:"url"`
	Id   string `json:"id"`
	Name string `json:"name"`
}
