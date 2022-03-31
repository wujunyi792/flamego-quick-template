package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/gin-template-new/internal/response/dto"
	"github.com/wujunyi792/gin-template-new/internal/service/jwtTokenGen"
	"net/http"
)

func JwtVerify(c *gin.Context) {
	var res dto.JsonResponse
	token := c.GetHeader("Authorization")
	if token != "" {
		entry, err := jwtTokenGen.ParseToken(token)
		if err == nil {
			c.Set("token", token)
			c.Set("id", entry.Info)
			c.Next()
			return
		} else {
			res.Code = 14005
			res.Message = fmt.Sprintf("%v", err)
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}
	}
	res.Code = 14005
	res.Message = "token not set, please add token in `Authorization` header"
	c.JSON(http.StatusForbidden, res)
	c.Abort()
	return
}
