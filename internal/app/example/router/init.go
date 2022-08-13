package router

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/internal/app/example/dto"
	"github.com/wujunyi792/gin-template-new/internal/app/example/handle"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
)

func AppExampleInit(e *flamego.Flame) {
	e.Group("/v1/example", func() {
		ExampleGroup(e)
	})
}

func ExampleGroup(e *flamego.Flame) {
	e.Get("", handle.HandelExampleGet)
	e.Post("", middleware.InjectRequest[dto.ExamplePost](), handle.HandelExamplePost)
}
