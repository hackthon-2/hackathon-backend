package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	database2 "hackthon/database/redis"
	"os"

	//"github.com/gofiber/fiber/v2"
	database "hackthon/database/mysql"
	"hackthon/model"
	"hackthon/util"
	"time"
)

var (
	ExistedUsername    = errors.New("username existed")
	RegistrationError  = errors.New("error when registered")
	NotFoundedUsername = errors.New("username not founded")
	InvalidPasswd      = errors.New("invalid password")
	NotMatched         = errors.New("no such email")
)

// Register 注册逻辑
func Register(input *model.RegisterInput) error {
	_, rows, _ := database.FindUserByUsername(input.Username)
	if rows != 0 {
		return ExistedUsername
	}
	cryptPasswd, err := util.HashPassword(input.Password)
	if err != nil {
		return err
	}
	input.Password = cryptPasswd
	var user model.User
	util.StructAssign(&user, input)
	err, rows = database.CreateUser(&user)
	if err != nil {
		return err
	}
	if rows != 1 {
		return RegistrationError
	}
	return nil
}

//Login 登陆逻辑
func Login(input *model.LoginInput) (string, error) {
	user, rows, err := database.FindUserByUsername(input.Username)
	if err != nil {
		return "", NotFoundedUsername
	}
	if rows != 1 {
		return "", NotFoundedUsername
	}
	match := util.CheckPasswordHash(input.Password, user.Password)
	if !match {
		return "", InvalidPasswd
	}
	var token string
	j := util.NewJWT()
	claim := util.Claims{
		Id:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "xzh",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token, err = j.CreateToken(claim)
	if err != nil {
		return "", err
	}
	err = database2.CreateTokenCache(user.Username,token)
	if err != nil {
		return "", err
	}
	return token, nil
}

// SendCode TODO:找回密码
func SendCode(input *model.FindPasswdInput) error {
	config := util.Mail{}
	username := os.Getenv("SMTP_USERNAME")
	passwd := os.Getenv("SMTP_PASSWD")
	config.SetMsg("smtp.qq.com", "465", username, passwd)
	//一个邮箱可能有多个用户名，要在验证code之后通过邮箱加载要选择的用户名进行密码重置
	data, err := database.FindUserByEmail(input.Email)
	if err != nil {
		return err
	}
	if len(data) < 1 {
		_ = config.SendTLS(input.Email, "重置密码邮件（您的邮箱可能已被别人所利用)", "经过查询，在我们的数据库中未找到关于您的邮箱的注册信息，但有人利用您的邮箱地址在我们的网站上进行重置密码操作。望知悉")
		return NotMatched
	}
	code := util.GenerateCode()
	err = database2.StoreCode(code)
	if err != nil {
		return err
	}
	err = config.SendTLS(input.Email, "重置密码邮件", input.Email+",您好！以下是您重置密码的链接：\r\n"+
		"https://www.onesnowwarrior.cn/findPasswd.html?code="+code+"&email="+input.Email+" (该链接十分钟内有效)\r\n"+
		"若此非本人操作，请忽视")
	return err
}
