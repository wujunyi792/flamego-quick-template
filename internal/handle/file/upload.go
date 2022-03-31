package file

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/gin-template-new/internal/response/dto"
	"github.com/wujunyi792/gin-template-new/internal/service/oss"
)

func HandleGetAliUploadToken(c *gin.Context) {
	resp := dto.JsonResponse{}
	resp.Clear()
	resp.Data = oss.GetPolicyToken()
	c.JSON(resp.Code/100, resp)
}
