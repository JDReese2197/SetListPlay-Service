package setlistFmModels

type User struct {
	UserId   string `json:"userId"`
	Fullname string `json:"fullname"`
	LastFm   string `json:"lastFm"`
	MySpace  string `json:"mySpace"`
	Twitter  string `json:"twitter"`
	Flickr   string `json:"flickr"`
	Website  string `json:"website"`
	About    string `json:"about"`
	Url      string `json:"url"`
}
