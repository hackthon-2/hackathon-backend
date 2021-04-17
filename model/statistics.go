package model
type Item struct {
	Item string `json:"item"`
	Times uint `json:"times"`
}
type Statistics struct {
	Header string `json:"header"`
	Items []Item `json:"todoItems"`
}
