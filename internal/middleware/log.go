package middleware

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/pkg/logx"
	"time"
)

func RequestLog() flamego.Handler {
	return func(c flamego.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request().Method

		// 请求路由
		reqUri := c.Request().RequestURI

		// 状态码
		statusCode := c.ResponseWriter().Status()

		// 请求IP
		clientIP := c.RemoteAddr()

		// 日志格式
		logx.Info.Printf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}
}
