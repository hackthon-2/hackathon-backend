package handler

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/constant"
	"hackthon/model"
	"hackthon/service"
	"hackthon/util"
	"log"
	"strconv"
	time2 "time"
)

// CreateDiary POST方式
func CreateDiary(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.DiaryInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if input.Time == "" || input.Question == "" {
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err := service.CreateDiary(claim.Id, &input)
	if err != nil {
		if err == service.ErrorWhenCreateDiary {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_301, constant.GetCodeText(constant.CODE_301))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

// UpdateDiary POST方式
func UpdateDiary(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.UpdateDiaryInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if input.Time == "" || input.Question == "" || input.ID == 0 {
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	var diary model.DiaryInput
	util.StructAssign(&diary, &input)
	err := service.UpdateDiary(claim.Id, input.ID, &diary)
	if err != nil {
		if err == service.ErrorWhenUpdateDiary {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_302, constant.GetCodeText(constant.CODE_302))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

// ListDiary GET方式
func ListDiary(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	time := c.Query("time", time2.Now().Format("2006-01-02"))
	data, err := service.ListDiary(claim.Id, time)
	if err != nil {
		if err == service.GetDiaryListError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_207, constant.GetCodeText(constant.CODE_207))
		}
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_207, constant.GetCodeText(constant.CODE_207))
	}
	return SuccessWithMessage(c, "获取成功", data)
}

// DeleteDiary DELETE方法
func DeleteDiary(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	diaryIDVal, err := strconv.Atoi(c.Query("diaryID"))
	if err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if diaryIDVal < 1 {
		log.Println("diary id out of range")
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err = service.DeleteDiary(claim.Id, uint(diaryIDVal))
	if err != nil {
		if err == service.DeleteDiaryError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_303, constant.GetCodeText(constant.CODE_303))
		}
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_303, constant.GetCodeText(constant.CODE_303))
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}
