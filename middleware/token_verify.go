package middleware

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
	database2 "hackthon/database/redis"
	"hackthon/handler"
	"hackthon/util"
	"log"
	"strings"
	"time"
)

func TokenVerify(c *fiber.Ctx) error {
	value := c.Get("Authorization")
	data := strings.Split(value, " ")
	if data[1] == "" {
		return handler.ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	token := data[1]
	log.Println("get token: " + token)
	j := util.NewJWT()
	claim, err := j.ParserToken(token)
	if err != nil {
		if err == util.TokenExpired {
			t, er := database2.FindTokenCache(token, time.Hour*4)
			if er != nil {
				return handler.ErrorWithMessage(c, constant.CODE_102, constant.GetCodeText(constant.CODE_102))
			}
			token, er = j.RefreshToken(token, t)
			if er != nil {
				return handler.ErrorWithMessage(c, constant.CODE_103, constant.GetCodeText(constant.CODE_103))
			}
			claim, _ = j.ParserToken(token)
			er = database2.CreateTokenCache(claim.Username, token)
			if er != nil {
				return handler.ErrorWithMessage(c, constant.CODE_103, constant.GetCodeText(constant.CODE_103))
			}
		}
	}
	c.Locals("claim", claim)
	return c.Next()
}
