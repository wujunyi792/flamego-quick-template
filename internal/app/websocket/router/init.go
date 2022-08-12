package router

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/internal/app/websocket/handle"
	"github.com/wujunyi792/gin-template-new/internal/app/websocket/service"
)

func AppWebsocketInit(e *flamego.Flame) {
	service.SocketInit()
	e.Group("/v1/websocket", func() {
		e.Get("/echo", handle.HandelWebsocketEcho)
	})
}
