package handle

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/internal/middleware"
	"github.com/wujunyi792/gin-template-new/internal/service/oss"
)

func HandelGetAliUploadToken(r flamego.Render) {
	middleware.HTTPSuccess(r, oss.GetPolicyToken())
}

func HandelAliUpLoad(r flamego.Render, c flamego.Context) {
	file, header, err := c.Request().FormFile("upload")
	if err != nil {
		middleware.InValidParam(r)
	} else {
		url := oss.UploadFileToOss(header.Filename, file)
		if url == "" {
			middleware.ServiceErr(r)
			return
		}
		middleware.HTTPSuccess(r, url)
	}
}
