package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
)

var registerTemplate = `
<html><body>
<div class="wrapper" style="margin: 20px auto 0; width: 500px; padding-top:16px; padding-bottom:10px;">
<div class="header clearfix"></div>
<br style="clear:both; height:0">
<div class="content" style="background: none repeat scroll 0 0 #FFFFFF; border: 1px solid #E9E9E9; margin: 2px 0 0; padding: 30px;">
<p>您好，您申请注册了 mkzilla 旗下 koala 平台的账号，请点击下面的链接输入密码完成注册。</p>
<a href="URL">
<div style="width: 40%;position: relative;padding:6px 15px 6px 15px; border-radius:18px; text-align:center;text-decoration:none;background-color:#2394f1;color:#ffffff;font-size:16px;margin:0;white-space:nowrap">
输入新密码完成注册
</div>
</a>
<p class="footer" style="border-top: 1px solid #DDDDDD; padding-top:6px; margin-top:25px; color:#838383;">© 2019 mkzilla koala&nbsp;&nbsp;|&nbsp;&nbsp;<a href="https://beacon.zhihu.com/c?c=e866ab50add2bd7fc8bb02ae7bc1c88d3d61e45d4b7b3b9d21c5637be1f412e7&amp;i=7bad544b48bc4209801abd52c7f5564b&amp;s=1&amp;t=1571729382&amp;u=http%3A%2F%2Fwww.zhihu.com%2Fl%2FIEswS&amp;utm_campaign=email_admin&amp;utm_source=zhihu-web" style="color:#838383;" rel="noopener" target="_blank">邮件通知设置</a></p>
</div>
</div></body></html>`

func SendMail(addr, from, to, subject, url string) error {
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

	body := strings.ReplaceAll(registerTemplate, "URL", url)

	data := fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", to, subject, body)
	buf := bytes.NewBufferString(data)
	_, err = buf.WriteTo(wc)
	return err
}
