package handler

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
	"hackthon/model"
	"hackthon/service"
	"log"
)

func Register(c *fiber.Ctx) error {
	var input model.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	e := service.Register(&input)
	if e != nil {
		log.Println(e.Error())
		if e == service.ExistedUsername {
			return ErrorWithMessage(c, constant.CODE_202, constant.GetCodeText(constant.CODE_202))
		}
		return StatusServerErrorWithMessage(c, constant.GetCodeText(constant.CODE_500))
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func Login(c *fiber.Ctx) error {
	var input model.LoginInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	token, e := service.Login(&input)
	if e != nil {
		log.Println(e.Error())
		if e == service.NotFoundedUsername {
			return ErrorWithMessage(c, constant.CODE_204, constant.GetCodeText(constant.CODE_204))
		}
		if e == service.InvalidPasswd {
			return ErrorWithMessage(c, constant.CODE_203, constant.GetCodeText(constant.CODE_203))
		}
		return StatusServerErrorWithMessage(c, constant.GetCodeText(constant.CODE_500))
	}
	return SuccessWithMessage(c, "请求成功", map[string]interface{}{
		"token": token,
	})
}
