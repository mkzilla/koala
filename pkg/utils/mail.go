package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
)

func SendMail(addr, from, to, subject, body string) error {
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	// 设置寄件人和收件人
	client.Mail(from)
	client.Rcpt(to)
	// 发送邮件主体。
	wc, err := client.Data()
	if err != nil {
		return err
	}
	defer wc.Close()

	data := fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", to, subject, body)
	buf := bytes.NewBufferString(data)
	_, err = buf.WriteTo(wc)
	return err
}

func SendMailWithPassword(hostname string, port int64, from, password, to, subject, body string) error {
	auth := smtp.PlainAuth(password, from, password, hostname)
	data := fmt.Sprintf("To: %s\r\nFrom: koala 官方<%s>\r\nSubject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", to, from, subject, body)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), auth, from, []string{to}, []byte(data))
	return err
}
