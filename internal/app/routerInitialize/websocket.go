package routerInitialize

import "github.com/wujunyi792/gin-template-new/internal/app/websocket/router"

func init() {
	routers = append(routers, router.AppWebsocketInit)
}
