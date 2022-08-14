package router

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/app/websocket/handle"
	"github.com/wujunyi792/flamego-quick-template/internal/middleware"
	"github.com/wujunyi792/flamego-quick-template/internal/websocket"
)

func AppWebsocketInit(e *flamego.Flame) {
	websocket.InitSocketManager("example")
	e.Group("/v1/websocket", func() {
		e.Use(middleware.InjectWebsocket("example"))
		e.Get("/echo", handle.HandelWebsocketEcho)
	})
}
