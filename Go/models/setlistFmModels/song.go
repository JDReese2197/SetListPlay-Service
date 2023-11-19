package setlistFmModels

type Song struct {
	Name  string `json:"name"`
	With  Artist `json:"with"`
	Cover Artist `json:"cover"`
	Info  string `json:"info"`
	Tape  bool   `json:"tape"`
}
