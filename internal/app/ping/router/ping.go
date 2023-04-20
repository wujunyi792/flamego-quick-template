package router

import (
	"errors"
	"github.com/flamego/binding"
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping/dto"
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping/handler"
	"github.com/wujunyi792/flamego-quick-template/internal/middleware/response"
)

func AppPingInit(e *flamego.Flame) {
	e.Get("/ping/v1", handler.HandleExampleGet)

	e.Get("/ping/v1/err", func(r flamego.Render) {
		response.HTTPFail(r, 500000, "test error", errors.New("this is err"))
	})

	e.Post("/ping/v1", binding.JSON(dto.ExamplePost{}), handler.HandlePing)
}
