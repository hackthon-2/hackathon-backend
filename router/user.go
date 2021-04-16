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
	user.Post("/diaryCreation", handler.CreateDiary)
	user.Post("/diaryUpdate", handler.UpdateDiary)
	user.Get("/diaryList", handler.ListDiary)
	user.Delete("/diaryDeletion", handler.DeleteDiary)
	user.Post("/todoCreation", handler.CreateTodo)
	user.Post("/todoUpdate", handler.UpdateTodo)
	user.Get("/todoList", handler.ListTodo)
	user.Delete("/todoDeletion", handler.DeleteTodo)
	user.Get("/statistics", handler.Statistics)
	user.Post("/avatar", handler.Uploader)
}
