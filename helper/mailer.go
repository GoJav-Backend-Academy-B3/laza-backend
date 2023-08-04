package helper

import (
	"os"
	"strconv"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type DataMail struct {
	Username  string
	Email     string
	Code      string
	EmailBody string
	Subject   string
}

func (res *DataMail) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Lazapedia <example@gmail.com>")
	m.SetHeader("To", res.Email)
	m.SetHeader("Subject", res.Subject)
	m.SetBody("text/html", res.EmailBody)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func Mail(data *DataMail) *DataMail {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Lazapedia",
			Link: os.Getenv("BASE_URL"),
			Logo: "https://cdn.discordapp.com/attachments/776427670979215363/1137076469877710888/412-4127373_datei-anzeigen-golang-gopher-china.jpeg",
		},
	}

	emailBody, _ := h.GenerateHTML(hermes.Email{
		Body: hermes.Body{
			Name: data.Username,
			Intros: []string{
				"Welcome to Lazapedia!",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please click the following button to verify your email. This code expires in 30 minutes.",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  data.Code,
					},
				},
			},
		},
	})

	return &DataMail{
		Username:  data.Username,
		Email:     data.Email,
		Code:      data.Code,
		EmailBody: emailBody,
		Subject:   data.Subject,
	}
}