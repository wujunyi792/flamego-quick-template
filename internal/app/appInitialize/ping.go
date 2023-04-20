package appInitialize

import (
	"github.com/wujunyi792/flamego-quick-template/internal/app/ping"
)

func init() {
	apps = append(apps, &ping.Ping{Name: "ping module"})
}
