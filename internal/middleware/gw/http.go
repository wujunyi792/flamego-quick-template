package gw

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
	"net/http"
	"time"
)

func RequestLog() flamego.Handler {
	return func(c flamego.Context, r *http.Request) {

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		logx.NameSpace("Request").Infow("request log", "path", c.Request().RequestURI, "method", c.Request().Method, "ip", c.RemoteAddr(), "status", c.ResponseWriter().Status(), "duration", time.Now().Sub(startTime))
	}
}
