package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"hackthon/config"
	"hackthon/database"
	"hackthon/router"
	"log"
)

func main() {
	config.Config()
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	database.ConnectDB()
	database.ConnectRedis()
	app.Use(cors.New(), recover.New())
	router.Init(app)
	defer database.DisconnectDB()
	defer database.DisconnectRedis()
	log.Fatal(app.ListenTLS(":8000", "./key/1_api.onesnowwarrior.cn_bundle.crt", "./key/2_api.onesnowwarrior.cn.key"))
}
