package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"hackthon/middleware"
)

func Init(app *fiber.App){
	api:=app.Group("/api",logger.New())
	api.Static("/","./static")
	auth:=api.Group("/auth")
	authInit(auth)
	user:=api.Group("/user",middleware.TokenVerify)
	userInit(user)
}