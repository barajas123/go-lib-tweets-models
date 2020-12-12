package tweetModels

type Ticker struct {
	Ticker   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Sector   string `json:"sector"`
	Industry string `json:"industry"`
}
