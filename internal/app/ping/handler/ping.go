package handler

import (
	"fmt"
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping/dto"
	"github.com/wujunyi792/flamego-quick-template/internal/middleware/response"
)

func HandleExampleGet(c flamego.Context, r flamego.Render) {
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
	response.HTTPSuccess(r, data)
}

func HandlePing(r flamego.Render, req dto.ExamplePost) {
	response.HTTPSuccess(r, req)
}
