package routerInitialize

import "github.com/flamego/flamego"

var (
	routers = make([]func(e *flamego.Flame), 0)
)

func ApiInit(r *flamego.Flame) {
	for _, router := range routers {
		router(r)
	}
}
