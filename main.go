package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hackthon/config"
	"hackthon/database"
	"hackthon/router"
	"log"
)

func main(){
	config.Config()
	app:=fiber.New(fiber.Config{
		Prefork: true,
	})
	database.ConnectDB()
	database.ConnectRedis()
	app.Use(cors.New())
	router.Init(app)
	defer database.DisconnectDB()
	defer database.DisconnectRedis()
	log.Fatal(app.Listen(":8000"))
}