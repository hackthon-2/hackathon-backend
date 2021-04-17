package model

type WatchInput struct {
	Content string `json:"content"`
	Time uint `json:"time"`
	FinishedTime uint `json:"finished_time"`
}
