package routerInitialize

import file "github.com/wujunyi792/flamego-quick-template/internal/app/file/router"

func init() {
	routers = append(routers, file.AppFileInit)
}
