package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/gin-template-new/internal/response/dto"
	"github.com/wujunyi792/gin-template-new/internal/router/v1/fileRouter"
	"github.com/wujunyi792/gin-template-new/internal/router/v1/websocketRouter"
)

func MainRouter(e *gin.Engine) {
	e.Any("", func(c *gin.Context) {
		res := dto.JsonResponse{}
		res.Clear()
		res.Data = struct {
			UA   string
			Host string
		}{
			UA:   c.Request.Header.Get("User-Agent"),
			Host: c.Request.Host,
		}
		c.JSON(res.Code/100, res)
	})
	fileRouter.InitFileRouter(e)
	websocketRouter.InitWebSocketRouter(e)
}
