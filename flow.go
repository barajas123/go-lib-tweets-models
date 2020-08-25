package tweetModels

type Flow struct {
	TwId       string  `json:"twId"`
	Ticker     string  `json:"ticker"`
	Strike     float32 `json:"strike"`
	PC         string  `json:"pc"`
	Quantity   int     `json:"quantity"`
	Price      float32 `json:"price"`
	Expiration string  `json:"expiration"`
	BidAsk     string  `json:"bidAsk"`
	RefPrice   float32 `json:"refPrice"`
	TimeStamp  string  `json:"createdAt"`
	IsSweep    bool    `json:"isSweep"`
	IsOpening  bool    `json:"isOpening"`
	RefVolume  float32 `json:"refVolume"`
}
