package service

import (
	"log"

	"gopkg.in/gomail.v2"
)

// Mail : Datas needed to send mail
type Mail struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       string
	Subject  string
	Body     string
}

// SendMail : Send mail using .env vars, subject and body
func SendMail(subject string, body string) error {
	host, port, user, pass, from, to, err := GetVarsMail()
	if err != nil {
		return err
	}

	mail := Mail{
		Host:     host,
		Port:     port,
		Username: user,
		Password: pass,
		From:     from,
		To:       to,
		Subject:  subject,
		Body:     body,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
