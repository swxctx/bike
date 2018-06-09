package mail

import (
	"fmt"

	"github.com/domego/gokits/log"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(bikeId int32) {
	m := gomail.NewMessage()
	// 发件人
	m.SetAddressHeader("From", "swxctx@sina.com", "单车告警")
	// 收件人
	m.SetHeader("To", m.FormatAddress("swxctx@sina.cn", "Bike"))
	// 主题
	m.SetHeader("Subject", "单车异常告警")
	// 发送的body体
	body := fmt.Sprintf("<h4>您好，系统检测到您的单车(%d)数据存在异常，请前往查看。<h4><a href =\"http://localhost:8893/#/bike_alarm\">点击查看详情</a>", bikeId)
	m.SetBody("text/html", body)

	// 发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer("smtp.sina.com", 25, "swxctx@sina.com", "6803161730")
	if err := d.DialAndSend(m); err != nil {
		log.Errorf("DialAndSend err->%v", err)
	}
	log.Infof("mail send success...")
}
