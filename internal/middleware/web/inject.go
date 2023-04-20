package web

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/websocket"
)

func InjectWebsocket(key string) flamego.Handler {
	return func(c flamego.Context) {
		c.Map(websocket.GetSocketManager(key))
	}
}
