package handler

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
	"hackthon/service"
	"hackthon/util"
	"log"
	"strings"
)

func Uploader(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	file, err := c.FormFile("document")
	if err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_401, constant.GetCodeText(constant.CODE_401))
	}
	val := strings.Split(file.Filename, ".")
	file.Filename = util.FileNameHash(val[0]) + "." + val[1]
	err = c.SaveFile(file, "./avatars/"+file.Filename)
	if err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_401, constant.GetCodeText(constant.CODE_401))
	}
	err = service.UploadAvatar(claim.Id, file.Filename)
	if err != nil {
		if err == service.UpdateProfileError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_401, constant.GetCodeText(constant.CODE_401))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}
