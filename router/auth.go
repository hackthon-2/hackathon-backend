package router

import (
	"github.com/gofiber/fiber/v2"
	"hackthon/handler"
)

func authInit(auth fiber.Router) {
	auth.Post("/registration", handler.Register)
	auth.Post("/login", handler.Login)
}
