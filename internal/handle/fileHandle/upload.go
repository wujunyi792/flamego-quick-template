package fileHandle

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

// HandleAliUpLoad 通过业务服务器中转文件至OSS 表单提交 字段名upload
func HandleAliUpLoad(c *gin.Context) {
	res := dto.JsonResponse{}
	res.Clear()
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		res.Code = 20008
		res.Message = err.Error()
	} else {
		url := oss.UploadFileToOss(header.Filename, file)
		if url == "" {
			res.Code = 50006
			res.Message = "上传失败"
		}
		res.Data = url
	}

	c.JSON(res.Code/100, res)

}
