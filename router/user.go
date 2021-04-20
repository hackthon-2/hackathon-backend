package router

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/handler"
	"hackthon/util"
)

func userInit(user fiber.Router) {
	user.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello, " + c.Locals("claim").(*util.Claims).Username)
	})
	//日记模块
	user.Post("/diaryCreation", handler.CreateDiary)
	user.Post("/diaryUpdate", handler.UpdateDiary)
	user.Get("/diaryList", handler.ListDiary)
	user.Get("/diary", handler.FindDiary)
	user.Delete("/diaryDeletion", handler.DeleteDiary)
	//待办模块
	user.Post("/todoCreation", handler.CreateTodo)
	user.Post("/todoUpdate", handler.UpdateTodo)
	user.Get("/todoList", handler.ListTodo)
	user.Get("/todo", handler.FindTodo)
	user.Delete("/todoDeletion", handler.DeleteTodo)
	//统计模块
	user.Get("/statistics", handler.Statistics)
	//用户模块
	user.Put("/avatar", handler.Uploader)
	user.Get("/profile", handler.ListProfile)
	user.Post("/profileUpdate", handler.UpdateProfile)
	//监督模块
	user.Post("/watchCreation", handler.CreateWatch)
	user.Post("/watchUpdate", handler.UpdateWatch)
	user.Get("/watch", handler.FindWatch)
	user.Delete("/watchDeletion", handler.DeleteWatch)
}
