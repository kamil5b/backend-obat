package utilities

import (
	"log"

	"gopkg.in/gomail.v2"
)

/*
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Akun bokep <zeehunter24@gmail.com>"
const CONFIG_AUTH_EMAIL = "zeehunter24@gmail.com"
const CONFIG_AUTH_PASSWORD = "AGEN&%*&!"
*/

func SendEmail(reciever, header, msg, sender, smtp_host, auth_email, auth_password string, port int) {
	/*
		CONFIG_SENDER_NAME := GoDotEnvVariable("CONFIG_SENDER_NAME")
		CONFIG_SMTP_HOST := GoDotEnvVariable("CONFIG_SMTP_HOST")
		CONFIG_AUTH_EMAIL := GoDotEnvVariable("CONFIG_AUTH_EMAIL")
		CONFIG_AUTH_PASSWORD := GoDotEnvVariable("CONFIG_AUTH_PASSWORD")

		CONFIG_SMTP_PORT, err := strconv.Atoi(GoDotEnvVariable("CONFIG_SMTP_PORT"))
		if err != nil {
			log.Fatal(err.Error())
		}
	*/
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", sender)
	mailer.SetHeader("To", reciever)
	mailer.SetHeader("Subject", header)
	mailer.SetBody("text/html", msg)

	dialer := gomail.NewDialer(
		smtp_host,
		port,
		auth_email,
		auth_password,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	//*/
	log.Println("Mail sent!")
}
