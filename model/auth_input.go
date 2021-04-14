package model

type LoginInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type RegisterInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}
type FindPasswdInput struct {
	Email string `form:"email"`
}

func (reg *RegisterInput) Init() {
	reg.Username = "guest"
	reg.Password = "000000"
	reg.Email = "guest@email.com"
}
