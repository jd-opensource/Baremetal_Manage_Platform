package util

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"html/template"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-gomail/gomail"
)

// SendMail use smtp to send mail
func SendMail(subject, msg string) error {
	m := gomail.NewMessage()
	from, _ := web.AppConfig.String("alert.mail.from")
	to, _ := web.AppConfig.Strings("alert.mail.to")
	s, _ := web.AppConfig.String("alert.mail.subject")
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", s+":"+subject)
	m.SetBody("text/html", msg)

	smtpHost, _ := web.AppConfig.String("alert.mail.smtp-host")
	username, _ := web.AppConfig.String("alert.mail.username")
	password, _ := web.AppConfig.String("alert.mail.password")

	d := gomail.NewPlainDialer(
		smtpHost,
		web.AppConfig.DefaultInt("alert.mail.smtp-port", 25),
		username,
		password)

	if d == nil {
		return errors.New("Initialize Dialer Error")
	}

	d.TLSConfig = &tls.Config{ServerName: d.Host, InsecureSkipVerify: true}

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("SendMail msg=%s : %s\n", msg, err.Error())
		return err
	}
	return nil
}

//SendMailByTpl send template mail
func SendMailByTpl(tplName, subject string, obj interface{}) error {
	tpl, errParse := template.New("Mail").Parse(tplName)
	if errParse != nil {
		fmt.Printf("Parse Template Error:%s\n", errParse.Error())
		return errParse
	}

	byteW := new(bytes.Buffer)
	err := tpl.Execute(byteW, obj)
	if err != nil {
		fmt.Printf("Execute Template Error:%s\n", err.Error())
		return err
	}

	errMail := SendMail(subject, byteW.String())
	if errMail != nil {
		fmt.Printf("Send Mail Error:%s\n", errMail.Error())
		return errMail
	}

	return nil
}
