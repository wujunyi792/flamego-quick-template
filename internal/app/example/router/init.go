package router

import "github.com/flamego/flamego"

func AppExampleInit(e *flamego.Flame) {
	e.Group("/v1/example", func() {
		ExampleGroup(e)
	})
}
