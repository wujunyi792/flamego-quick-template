package routerInitialize

import "github.com/wujunyi792/flamego-quick-template/internal/app/websocket/router"

func init() {
	routers = append(routers, router.AppWebsocketInit)
}
