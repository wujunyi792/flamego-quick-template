package router

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/app/file/handler"
)

func AliGroup(e *flamego.Flame) {
	e.Group("/ali", func() {
		e.Get("/token", handler.HandleGetAliUploadToken)
		e.Post("/upload", handler.HandleAliUpLoad)
	})
}
