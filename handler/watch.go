package handler

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
	"hackthon/model"
	"hackthon/service"
	"hackthon/util"
	"log"
)

func CreateWatch(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.WatchInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err := service.CreateWatch(claim.Username, claim.Id, &input)
	if err != nil {
		if err == service.ErrorWhenCreatingWatch {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_308, constant.GetCodeText(constant.CODE_308))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func UpdateWatch(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.UpdateWatchInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err := service.UpdateWatch(input.ID, claim.Id, &input)
	if err != nil {
		if err == service.ErrorWhenUpdatingWatch {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_309, constant.GetCodeText(constant.CODE_309))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func FindWatch(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	data, err := service.FindWatch(claim.Id)
	if err != nil {
		if err == service.FindWatchError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_310, constant.GetCodeText(constant.CODE_310))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "获取成功", data)
}

func DeleteWatch(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	err := service.DeleteWatch(claim.Id)
	if err != nil {
		if err == service.ErrorWhenDeletingWatch {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_311, constant.GetCodeText(constant.CODE_311))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}
