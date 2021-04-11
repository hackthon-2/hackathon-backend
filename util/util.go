package util

import (
	"crypto/rsa"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"reflect"
	"time"
)

var (
	TokenExpired     = errors.New("Token 已经过期")
	TokenNotValidYet = errors.New("Token 未激活")
	TokenMalformed   = errors.New("Token 错误")
	TokenInvalid     = errors.New("Token 无效")
)

type JWT struct {
	SigningKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}
type Claims struct {
	Id       uint
	Username string
	jwt.StandardClaims
}

// NewJWT 新建一个JWT的结构体对象，并把读取的公私钥信息进行反序列化供程序使用
func NewJWT() *JWT {
	SigningBytes, err := ioutil.ReadFile("./key/pkcs8_rsa_private_key.pem")
	if err != nil {
		log.Println(err.Error())
		log.Println("读入私钥失败")
		return nil
	}
	PublicBytes, err := ioutil.ReadFile("./key/rsa_public_key.pem")
	if err != nil {
		log.Println("读入公钥失败")
		return nil
	}
	SigningKey, err := jwt.ParseRSAPrivateKeyFromPEM(SigningBytes)
	if err != nil {
		log.Println("转换私钥失败")
		return nil
	}
	PublicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicBytes)
	if err != nil {
		log.Println("转换公钥失败")
		return nil
	}
	return &JWT{
		SigningKey: SigningKey,
		PublicKey:  PublicKey,
	}
}
func (j *JWT) CreateToken(c Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	return token.SignedString(j.SigningKey)
}
func (j *JWT) ParserToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c, nil
	}
	return nil, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string, t time.Duration) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Now().Add(t)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err != nil {
		return "", err
	}
	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		c.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*c)
	}
	return "", TokenInvalid
}
/*
StructAssign
binding type interface 要修改的结构体
value type interface 有数据的结构体
对两个结构体相同字段进行赋值操作
*/
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}