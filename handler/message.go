package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"hackthon/service"
	"hackthon/util"
)

func Message(hub *service.Hub) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		claim := c.Locals("claim").(*util.Claims)
		client := &service.Client{
			Hub:      hub,
			Conn:     c,
			Username: claim.Username,
			Send:     make(chan []byte, 256),
		}
		client.Hub.Register <- client
		go client.WritePump()
		client.ReadPump()
	}, websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
}
