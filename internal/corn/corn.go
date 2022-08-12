package corn

import (
	"github.com/robfig/cron"
	"github.com/wujunyi792/gin-template-new/internal/logx"
)

func init() {
	c := cron.New()
	err := c.AddFunc("0 0/10 * * * *", func() {})
	if err != nil {
		logx.Error.Fatalln(err)
	}
	c.Start()
	logx.Info.Println("corn init SUCCESS ")
}
