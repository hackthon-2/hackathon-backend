package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	database2 "hackthon/database/redis"

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
	err = database2.CreateTokenCache(token)
	if err != nil {
		return "", err
	}
	return token, nil
}
