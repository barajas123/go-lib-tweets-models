package tweetModels

import (
	"fmt"
	"strings"
)

type FlowRaw struct {
	TwId      string `json:"twid"`
	Text      string `json:"text"`
	TimeStamp string `json:"createdAt"`
}

// Classify is used to tag the TwID with the type of Tweet that it is (S,B,C)
func (f *FlowRaw) Classify() {
	f.Text = strings.ToLower(f.Text)
	if strings.Contains(f.Text, "sweep") {
		f.TwId = fmt.Sprintf("%sS", f.TwId)
		return
	}
	if strings.Contains(f.Text, "block") {
		f.TwId = fmt.Sprintf("%sB", f.TwId)
	} else {
		f.TwId = fmt.Sprintf("%sC", f.TwId)
	}
}
