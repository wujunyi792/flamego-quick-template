package appInitialize

import (
	"github.com/wujunyi792/flamego-quick-template/internal/app"
)

var (
	apps = make([]app.Module, 0)
)

func GetApps() []app.Module {
	return apps
}
