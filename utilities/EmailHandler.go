package utilities

import (
	"log"
	// "gopkg.in/gomail.v2"
)

/*
const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Akun bokep <zeehunter24@gmail.com>"
const CONFIG_AUTH_EMAIL = "zeehunter24@gmail.com"
const CONFIG_AUTH_PASSWORD = "AGEN&%*&!"
*/

func SendEmail(reciever, header, msg string) {
	/*
		CONFIG_SENDER_NAME := GoDotEnvVariable("CONFIG_SENDER_NAME")
		CONFIG_SMTP_HOST := GoDotEnvVariable("CONFIG_SMTP_HOST")
		CONFIG_AUTH_EMAIL := GoDotEnvVariable("CONFIG_AUTH_EMAIL")
		CONFIG_AUTH_PASSWORD := GoDotEnvVariable("CONFIG_AUTH_PASSWORD")

		CONFIG_SMTP_PORT, err := strconv.Atoi(GoDotEnvVariable("CONFIG_SMTP_PORT"))
		if err != nil {
			log.Fatal(err.Error())
		}

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

		err = dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}
	*/
	log.Println("Mail sent!")
}
