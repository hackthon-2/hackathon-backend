package model

type UpdateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
