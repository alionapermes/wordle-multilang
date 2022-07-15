package main

type WordOfDay struct {
	Id   int64  `json:"id"`
	Text string `json:"text"`
	Lang string `json:"lang"`
	Next int64  `json:"next"`
}

type Payload struct {
	Word string `json:"word"`
	Lang string `json:"lang"`
	Next int64  `json:"next"`
}
