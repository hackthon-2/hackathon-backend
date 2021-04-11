package model
type LoginInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type RegisterInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Sex      uint   `form:"sex"`
}