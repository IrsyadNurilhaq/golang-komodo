package utils

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject, uuid string) error {
	body := `<b> Please confirmation your account </b> : ` + viper.GetString("app.api_url") + `v1/confirmation?code=` + uuid
	var (
		CONFIG_SMTP_HOST = viper.GetString("email.SMTP_HOST")
		CONFIG_SMTP_PORT = viper.GetInt("email.SMTP_PORT")
		CONFIG_EMAIL     = viper.GetString("email.EMAIL")
		CONFIG_PASSWORD  = viper.GetString("email.PASSWORD")
	)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_EMAIL)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Confirmation Email")
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_EMAIL,
		CONFIG_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	return nil
}
