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
	//var rawJson string = `{"header":"daskljdakslf","todoItems":[{"id":"faskjfkaf","item":"图书馆","isComplete":"false"},{"id":"faskjfkafgdas","item":"图书","isComplete":"true"},{"id":"ffsaf","item":"书","isComplete":"false"}]}`
	//
	//var todo Todo
	//_=json.Unmarshal([]byte(rawJson),&todo)
	//fmt.Println(todo)

}
