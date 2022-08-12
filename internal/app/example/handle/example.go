package handle

import (
	"fmt"
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/config"
	"github.com/wujunyi792/gin-template-new/internal/app/example/dto"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
)

func HandelExampleGet(c flamego.Context, r flamego.Render) {
	data := struct {
		UA         string
		Host       string
		Method     string
		Proto      string
		RemoteAddr string
		Message    string
	}{
		UA:         c.Request().UserAgent(),
		Host:       c.Request().Host,
		Method:     c.Request().Method,
		Proto:      c.Request().Proto,
		RemoteAddr: c.Request().RemoteAddr,
		Message:    fmt.Sprintf("Welcome to %s, version %s.", config.GetConfig().ProgramName, config.GetConfig().VERSION),
	}
	middleware.HTTPSuccess(r, data)
}

func HandelExamplePost(r flamego.Render, req dto.ExamplePost) {
	middleware.HTTPSuccess(r, req.Msg)
}
