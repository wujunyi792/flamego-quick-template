package corn

import (
	"github.com/robfig/cron"
	"github.com/wujunyi792/gin-template-new/internal/loging"
)

func init() {
	c := cron.New()
	err := c.AddFunc("0 0/10 * * * *", func() {})
	if err != nil {
		loging.Error.Fatalln(err)
	}
	c.Start()
	loging.Info.Println("corn init SUCCESS ")
}
