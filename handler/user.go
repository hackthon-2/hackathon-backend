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

func UpdateProfile(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err := service.UpdateProfile(claim.Id, &input)
	if err != nil {
		if err == service.UpdateProfileError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_402, constant.GetCodeText(constant.CODE_402))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}
func ListProfile(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	user, err := service.ListProfile(claim.Id)
	if err != nil {
		if err == service.ListProfileError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_403, constant.GetCodeText(constant.CODE_403))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "获取成功", user)
}

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
		return StatusServerErrorWithMessage(c, err.Error())
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
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func CreateTodo(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.TodoInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if input.Time == "" || input.Header == "" {
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err := service.CreateTodo(claim.Id, &input)
	if err != nil {
		if err == service.CreateTodoError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_304, constant.GetCodeText(constant.CODE_304))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func UpdateTodo(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	var input model.UpdateTodoInput
	if err := c.BodyParser(&input); err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if input.Time == "" || input.Header == "" || input.ID == 0 {
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	var todo model.TodoInput
	util.StructAssign(&todo, &input)
	err := service.UpdateTodo(claim.Id, input.ID, &todo)
	if err != nil {
		if err == service.UpdateTodoError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_305, constant.GetCodeText(constant.CODE_305))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func ListTodo(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	time := c.Query("time", time2.Now().Format("2006-01-02"))
	data, err := service.ListTodo(claim.Id, time)
	if err != nil {
		if err == service.GetTodoListError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_207, constant.GetCodeText(constant.CODE_207))
		}
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "获取成功", data)
}

func DeleteTodo(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	todoIDVal, err := strconv.Atoi(c.Query("todoID"))
	if err != nil {
		log.Println(err.Error())
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	if todoIDVal < 1 {
		log.Println("todo id out of range")
		return ErrorWithMessage(c, constant.CODE_100, constant.GetCodeText(constant.CODE_100))
	}
	err = service.DeleteTodo(claim.Id, uint(todoIDVal))
	if err != nil {
		if err == service.DeleteTodoError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_306, constant.GetCodeText(constant.CODE_306))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "请求成功", [0]int{})
}

func Statistics(c *fiber.Ctx) error {
	claim := c.Locals("claim").(*util.Claims)
	time := c.Query("time", time2.Now().Add(-time2.Hour*168).Format("2006-01-02"))
	data, err := service.Statistics(claim.Id, time)
	if err != nil {
		if err == service.GenerateStatisticsError {
			log.Println(err.Error())
			return ErrorWithMessage(c, constant.CODE_307, constant.GetCodeText(constant.CODE_307))
		}
		log.Println(err.Error())
		return StatusServerErrorWithMessage(c, err.Error())
	}
	return SuccessWithMessage(c, "获取成功", data)
}
