package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App){
	api:=app.Group("/api",logger.New())
	api.Static("/","./static")
	auth:=api.Group("/auth")
	authInit(auth)
	/*
	TODO:搞个auth.go把auth的路由信息在authInit中初始化
	 */
}