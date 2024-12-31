package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"
	"time"

	"github.com/go-gomail/gomail"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var from string //发件人

var d *gomail.Dialer

func InitEmailConfig(c EmailConfig) error {
	d = gomail.NewDialer(c.Host, c.Port, c.Username, c.Password)
	from = c.From
	return nil
}

// SendMail use smtp to send mail
func SendMail(subject, msg string, receivers []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", receivers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	if d == nil {
		return errors.New("Initialize Dialer Error")
	}

	d.TLSConfig = &tls.Config{ServerName: d.Host, InsecureSkipVerify: true}

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(time.Now(), "SendMail msg=", msg, err.Error())
		return err
	}
	return nil
}

//本机对smtp.jd.local:25邮件服务没有访问权限，所以邮件投递能否成功没有测试
//SendMailByTpl send template mail
func SendMailByTpl(tplName, subject string, obj interface{}, receivers []string) error {
	tpl, errParse := template.New("Mail").Parse(tplName)
	if errParse != nil {
		fmt.Printf("Parse Template Error:%s\n", errParse.Error())
		return errParse
	}

	byteW := new(bytes.Buffer)
	err := tpl.Execute(byteW, obj)
	if err != nil {
		fmt.Println(time.Now(), "Execute Template Error:", err.Error())
		return err
	}

	errMail := SendMail(subject, byteW.String(), receivers)
	if errMail != nil {
		fmt.Println(time.Now(), "Send Mail Error:", errMail.Error())
		return errMail
	}

	return nil
}
