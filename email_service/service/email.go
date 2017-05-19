package service

import (
	"fmt"
	"model"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
)

var Cfg = beego.AppConfig

type EmailService struct {
}

func (this *EmailService) Send(emails []*model.Email) error {
	var err error

	count := len(emails)
	chs := make(chan error, count)
	for i := 0; i < count; i++ {
		// check to
		beego.Debug("->check target email address")
		if emails[i].To == "" {
			err = fmt.Errorf("%s", "please input target email address")
			return err
		}

		// check body
		beego.Debug("->check email body")
		if emails[i].Body == "" {
			err = fmt.Errorf("%s", "please input email body")
			return err
		}

		// send email
		systemEmailUser := Cfg.String("system_email_user")
		systemEmailPassword := Cfg.String("system_email_password")
		systemEmailHost := Cfg.String("system_email_host")

		beego.Debug("->send email")
		beego.Debug("from:", systemEmailUser)
		beego.Debug("host:", systemEmailHost)
		beego.Debug("to:", emails[i].To)
		beego.Debug("->send email to", emails[i].To)

		go this.sendMail(systemEmailUser, systemEmailPassword, systemEmailHost, emails[i].To, emails[i].Subject, emails[i].Body, "html", chs)
	}

	for i := 0; i < count; i++ {
		v, _ := <-chs
		fmt.Println("return :", v)
	}

	close(chs)

	return err
}

func (this *EmailService) sendMail(user, password, host, to, subject, body, mailtype string, ch chan error) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	fmt.Println(err)
	//return err
	// ch <- "send mail success:" + to

	ch <- err
}

func (this *EmailService) SendEmail(user, password, host, to, subject, body, mailtype string) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	fmt.Println(err)
}

// func main() {
// 	user := "suengguan@126.com"
// 	password := "Shengguan0285"
// 	host := "smtp.126.com:25"
// 	//to := "suengguan@126.com;qsg@corex-tek.com"

// 	subject := "登录提醒"

// 	//	body := `
// 	//    <html>
// 	//    <body>
// 	//    <h3>
// 	//	"【科思世纪官网登录提醒】欢迎您于 **日期 登录PME系统。特此提醒。【科思世纪】"
// 	//	</h3>
// 	//    </body>
// 	//    </html>
// 	//    `
// 	//time.Now().Format()
// 	currentTime := time.Now().Format("2006-01-02 15:04:05")
// 	body := "【科思世纪官网登录提醒】欢迎您于 " + currentTime + " 登录PME系统。特此提醒。【科思世纪】"

// 	fmt.Println(body)

// 	var toList []string
// 	to1 := "suengguan@126.com"
// 	toList = append(toList, to1)
// 	to2 := "qsg@corex-tek.com"
// 	toList = append(toList, to2)

// 	count := len(toList)
// 	chs := make(chan string, count)
// 	for i := 0; i < count; i++ {
// 		go SendMail(user, password, host, toList[i], subject, body, "html", chs)
// 	}

// 	for i := 0; i < count; i++ {
// 		v, _ := <-chs
// 		fmt.Println("return :", v)
// 	}

// 	close(chs)

// 	//	fmt.Println("send email")
// 	//	err := SendMail(user, password, host, to, subject, body, "html")
// 	//	if err != nil {
// 	//		fmt.Println("send mail error!")
// 	//		fmt.Println(err)
// 	//	} else {
// 	//		fmt.Println("send mail success!")
// 	//	}

// }
