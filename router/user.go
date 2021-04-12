package router

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/util"
)

func userInit(user fiber.Router) {
	user.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello, " + c.Locals("claim").(*util.Claims).Username)
	})
}
