package mail

import (
	"github.com/qiniu/log"

	gomail "gopkg.in/gomail.v2"
)

func SendMail() {
	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", "swxctx@sina.com", "单车告警")
	// 收件人
	m.SetHeader("To", m.FormatAddress("swxctx@sina.cn", "go-mail"))
	// 主题
	m.SetHeader("Subject", "单车数据异常")
	// 发送的body体
	m.SetBody("text/html", "<h4>单车状态异常<h4><a href =\"http://localhost:8893/#/bike_alarm\">点击查看详情</a>")

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer("smtp.sina.com", 25, "swxctx@sina.com", "****")
	if err := d.DialAndSend(m); err != nil {
		log.Errorf("DialAndSend err->%v", err)
	}
	log.Infof("mail send success...")
}
