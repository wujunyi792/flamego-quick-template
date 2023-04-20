package corn

import (
	"github.com/robfig/cron"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
)

var (
	log = logx.NameSpace("corn")
)

func init() {
	c := cron.New()
	err := c.AddFunc("0 0/10 * * * *", func() {})
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
	log.Infoln("corn routerInitialize SUCCESS ")
}
