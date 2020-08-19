package tweetModels

type Flow struct{
  TwId string `json:"twId"`
  Ticker string `json:"ticker"`
  Strike double `json:"strike"`
  PC string `json:"pc"`
  Quantity int `json:"quantity"`
  Price double `json:"price"`
  Expiration string `json:"expiration"`
  BidAsk string `json:"bidAsk"`
  RefPrice double `json:"refPrice"`
  TimeStamp string `json:"createdAt"`
  Aggressiveness string `json:"aggressiveness"`
}
