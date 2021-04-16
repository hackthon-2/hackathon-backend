package model

type ToDoItem struct {
	ID         string `json:"id"`
	Item       string `json:"item"`
	IsComplete bool   `json:"isComplete"`
}
type TodoResponse struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	Header    string     `json:"header"`
	ToDoItems []ToDoItem `json:"todoItems"`
	Time      string     `json:"time"`
}
type TodoInput struct {
	Header    string     `json:"header"`
	ToDoItems []ToDoItem `json:"todoItems"`
	Time      string     `json:"time"`
}
type UpdateTodoInput struct {
	ID        uint       `json:"id"`
	Header    string     `json:"header"`
	ToDoItems []ToDoItem `json:"todoItems"`
	Time      string     `json:"time"`
}
