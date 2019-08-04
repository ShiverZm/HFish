package send

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"strconv"
	"time"
)

func SendMail(mailTo []string, subject string, body string, config []string) error {
	port, _ := strconv.Atoi(config[1])
	m := gomail.NewMessage()

	m.SetHeader("From", "<"+config[2]+">")
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(config[0], port, config[2], config[3])

	err := d.DialAndSend(m)

	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, time.Now().Format("2006-01-02 15:04:05")+" 发送邮件通知失败 ", err)
	} else {
		fmt.Fprintln(gin.DefaultWriter, time.Now().Format("2006-01-02 15:04:05")+" 发送邮件通知成功")
	}

	return err
}
