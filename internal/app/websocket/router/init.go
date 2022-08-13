package router

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/internal/app/websocket/handle"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
	"github.com/wujunyi792/gin-template-new/internal/websocket"
)

func AppWebsocketInit(e *flamego.Flame) {
	websocket.InitSocketManager("example")
	e.Group("/v1/websocket", func() {
		e.Use(middleware.InjectWebsocket("example"))
		e.Get("/echo", handle.HandelWebsocketEcho)
	})
}
