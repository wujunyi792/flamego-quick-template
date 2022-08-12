package router

import "github.com/flamego/flamego"

func AppFileInit(e *flamego.Flame) {
	e.Group("/v1/file", func() {
		AliGroup(e)
	})
}
