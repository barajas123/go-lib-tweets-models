package tweetModels

type News struct {
	TwId      string `json:"twId"`
	User      string `json:"username"`
	Ticker    string `json:"ticker"`
	Text      string `json:"text"`
	TimeStamp string `json:"created_at"`
}
