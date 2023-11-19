package setlistFmModels

type SetlistResult struct {
	Setlist      []Setlist `json:"setlist"`
	Total        int       `json:"total"`
	Page         int       `json:"page"`
	ItemsPerPage int       `json:"itemsPerPage"`
}
