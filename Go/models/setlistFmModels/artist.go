package setlistFmModels

//	to print use:	 fmt.Printf("%+v\n", yourStruct)

type Artist struct {
	Mbid           string `json:"mbid"`
	Tmid           int    `json:"tmid"`
	Name           string `json:"name"`
	SortName       string `json:"sortName"`
	Disambiguation string `json:"disambiguation"`
	Url            string `json:"url"`
}
