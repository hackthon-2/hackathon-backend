package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
	"time"
)

func WebsocketInit(ws fiber.Router) {
	ws.Get("/message", websocket.New(func(ctx *websocket.Conn) {
		var (
			mt  int
			msg []byte
			err error
		)
		ctx.Subprotocol()
		for {
			if mt, msg, err = ctx.ReadMessage(); err != nil {
				log.Println("read: ", err)
				break
			}
			log.Printf("recv: %v,%s", mt, msg)
			if err = ctx.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}, websocket.Config{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 20 * time.Second,
	}))
}
