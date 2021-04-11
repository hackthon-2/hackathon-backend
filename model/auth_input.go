package model

type LoginInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type RegisterInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Sex      uint   `form:"sex"` //0是男性，1是女性
}

func (reg *RegisterInput) Init() {
	reg.Username = "guest"
	reg.Password = "000000"
	reg.Email = "guest@email.com"
	reg.Sex = 0
}
