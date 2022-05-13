package utilities

import (
	"log"

	"gopkg.in/gomail.v2"
)

var (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "Nama Pengirim <inipengirim@gmail.com>"
	CONFIG_AUTH_EMAIL    = "inipengirim@gmail.com"
	CONFIG_AUTH_PASSWORD = "INIPASSWORD!"
)

func SetEmailConfig(smtp_host string, smtp_port int, sender_name, auth_email, auth_password string) {
	CONFIG_SMTP_HOST = smtp_host
	CONFIG_SMTP_PORT = smtp_port
	CONFIG_SENDER_NAME = sender_name
	CONFIG_AUTH_EMAIL = auth_email
	CONFIG_AUTH_PASSWORD = auth_password
}

func SendEmail(reciever, header, msg string) {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", reciever)
	mailer.SetHeader("Subject", header)
	mailer.SetBody("text/html", msg)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
