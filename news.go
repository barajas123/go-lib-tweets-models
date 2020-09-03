package tweetModels

type News struct {
	TwId      string `json:"twId"`
	User      string `json:"username"`
	Ticker    string `json:"ticker"`
	Headline  string `json:"headline"`
	TimeStamp string `json:"created_at"`
}
