package model

type WatchInput struct {
	Content      string `json:"content"`
	Time         uint   `json:"time"`
	FinishedTime uint   `json:"finished_time"`
}
type UpdateWatchInput struct {
	ID           uint   `json:"watch_id"`
	Watcher      string `json:"watcher"`
	WatcherID    uint   `json:"watcher_id"`
	FinishedTime uint   `json:"finished_time"`
}
