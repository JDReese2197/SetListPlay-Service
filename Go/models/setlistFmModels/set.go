package setlistFmModels

type Set struct {
	Name   string `json:"name"`
	Encore int    `json:"encore"`
	Song   []Song `json:"song"`
}
