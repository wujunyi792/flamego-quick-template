package middleware

import (
	"encoding/json"
	"github.com/flamego/flamego"
)

type resp struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func InjectRequest[T any]() flamego.Handler {
	var req T
	return func(r flamego.Render, c flamego.Context) {
		body, _ := c.Request().Body().Bytes()
		if err := json.Unmarshal(body, &req); err != nil {
			InValidParam(r)
			return
		}
		c.Map(req)
	}
}

func http(r flamego.Render, code int, msg string, data interface{}, count ...int) {
	co := 0
	if len(count) > 0 {
		co = count[0]
	}
	r.JSON(code/100, &resp{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: co,
	})
}

// HTTPSuccess 成功返回
func HTTPSuccess(r flamego.Render, data interface{}, count ...int) {
	http(r, 20000, "success", data, count...)
}

func HTTPFail(r flamego.Render, code int, msg string, count ...int) {
	http(r, code, msg, nil, count...)
}

func UnAuthorization(r flamego.Render) {
	HTTPFail(r, 40100, "unAuthorized")
}

func InValidParam(r flamego.Render) {
	HTTPFail(r, 40200, "invalid params")
}

func ServiceErr(r flamego.Render) {
	HTTPFail(r, 500, "service down")
}
