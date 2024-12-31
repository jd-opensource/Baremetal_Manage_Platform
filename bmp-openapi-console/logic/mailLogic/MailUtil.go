package mailLogic

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/mailDao"
	log "coding.jd.com/aidc-bmp/bmp_log"

	"github.com/go-gomail/gomail"
)

// SendMail use smtp to send mail
func SendMail(logger *log.Logger, subject, msg string, receivers string) error {

	mailEntity, err := mailDao.GetMailMessage(logger)
	if err != nil {
		logger.Warnf("SendMail.GetMail error:%s", err.Error())
		return err
	}
	tos := mailEntity.To
	from := mailEntity.AdminAddr
	smtpHost := mailEntity.ServerAddr
	smtpPort, err := strconv.Atoi(mailEntity.ServerPort)
	if err != nil {
		logger.Warnf("SendMail.ServerPort invalid, port:%s", mailEntity.ServerPort)
		return err
	}
	username := strings.Split(from, "@")[0]
	password := mailEntity.AdminPass

	m := gomail.NewMessage()

	to := strings.Split(tos, ",")
	if receivers != "" { //如果设置了收件人，就用设置的收件人，否则用系统收件人
		to = strings.Split(receivers, ",")
	}
	s := BaseSubjest
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", s+":"+subject)
	m.SetBody("text/html", msg)

	d := gomail.NewPlainDialer(
		smtpHost,
		smtpPort,
		username,
		password)

	if d == nil {
		return errors.New("Initialize Dialer Error")
	}

	d.TLSConfig = &tls.Config{ServerName: d.Host, InsecureSkipVerify: true}

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		logger.Warnf("SendMail msg=%s : %s\n", msg, err.Error())
		return err
	}
	return nil
}

//SendMailByTpl send template mail
func SendMailByTpl(logger *log.Logger, tplName, subject string, obj interface{}, receivers string) error {
	tpl, errParse := template.New("Mail").Parse(tplName)
	if errParse != nil {
		logger.Warnf("Parse Template Error:%s\n", errParse.Error())
		return errParse
	}

	byteW := new(bytes.Buffer)
	err := tpl.Execute(byteW, obj)
	if err != nil {
		logger.Warnf("Execute Template Error:%s\n", err.Error())
		return err
	}

	errMail := SendMail(logger, subject, byteW.String(), receivers)
	if errMail != nil {
		logger.Warnf("Send Mail Error:%s\n", errMail.Error())
		return errMail
	}

	return nil
}
