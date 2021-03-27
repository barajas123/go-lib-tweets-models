package tweetModels

type Ticker struct {
	Ticker      string `json:"symbol"`
	Name        string `json:"name"`
	Sector      string `json:"sector"`
	Industry    string `json:"industry"`
	Description string `json:"description"`
}
