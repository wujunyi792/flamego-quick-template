package handle

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/flamego-quick-template/internal/middleware"
	"github.com/wujunyi792/flamego-quick-template/internal/service/oss"
)

func HandleGetAliUploadToken(r flamego.Render) {
	middleware.HTTPSuccess(r, oss.GetPolicyToken())
}

func HandleAliUpLoad(r flamego.Render, c flamego.Context) {
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
