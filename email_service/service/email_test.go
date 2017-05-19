package service

import (
	//"common/model"
	"testing"
	//"time"
)

func Test_EmailService_SendEmail(t *testing.T) {
	systemEmailUser := "corex_tek@126.com"
	systemEmailPassword := "corex123"
	systemEmailHost := "smtp.126.com:25"

	t.Log("->send email")
	t.Log("from:", systemEmailUser)
	t.Log("host:", systemEmailHost)
	//t.Log("to:", emails[i].To)
	//t.Log("->send email to", emails[i].To)

	subject := "使用Golang发送邮件"
	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`

	var svc EmailService
	svc.SendEmail(systemEmailUser, systemEmailPassword, systemEmailHost, "suengguan@126.com", subject, body, "html")

	//t.Log(err)
}
