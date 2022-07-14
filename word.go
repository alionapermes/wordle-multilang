package main

type WordOfDay struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
	Lang string `json:"lang"`
	Day  string `json:"day"`
}

type Payload struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
	next int64  `json:"next"`
}
