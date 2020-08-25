package tweetModels

type Tweet struct {
	TwId      string `json:"id_str"`
	User      string `json:"user.screen_name"`
	Text      string `json:"text"`
	TimeStamp string `json:"created_at"`
}
