package mail

import (
	"fmt"
	"github.com/wujunyi792/flamego-quick-template/config"
	"github.com/wujunyi792/flamego-quick-template/internal/core/logx"
	"gopkg.in/gomail.v2"
	"time"
)

var (
	log = logx.NameSpace("mail")
)

var mailConfig struct {
	Instance *gomail.Dialer
}

func InitMail() {

	conf := &config.GetConfig().Mail

	mailConfig.Instance = gomail.NewDialer(
		conf.SMTP,
		conf.PORT,
		conf.ACCOUNT,
		conf.PASSWORD,
	)

	err := SendMail(conf.ACCOUNT, config.GetConfig().ProgramName+` Golang Program routerInitialize`, fmt.Sprintf("Name: %s\nVERSION: %s\nAuthor: %s\nTime: %s", config.GetConfig().ProgramName, config.GetConfig().VERSION, config.GetConfig().AUTHOR, time.Now().Format("2006-01-02 15:04:05")))

	if err != nil {
		log.Fatalln(err)
	}
	log.Infoln("mail routerInitialize SUCCESS ")
}

func SendMail(to, title, content string) error {
	conf := &config.GetConfig().Mail
	m := gomail.NewMessage()
	m.SetHeader("From", "Golang Program Manager"+"<"+conf.ACCOUNT+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)

	return mailConfig.Instance.DialAndSend(m)
}
