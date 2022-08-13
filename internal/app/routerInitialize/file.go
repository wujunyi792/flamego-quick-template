package routerInitialize

import file "github.com/wujunyi792/gin-template-new/internal/app/file/router"

func init() {
	routers = append(routers, file.AppFileInit)
}
