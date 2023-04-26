package mailer

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	client   *sendgrid.Client
	fromName string
	fromMail string
}

func NewSendGrid(apiKey, fromName, fromMail string) Mailer {
	if "apiKey" == "" {
		log.Fatal().Msg("missing sendgrid api key")
	}
	client := sendgrid.NewSendClient(apiKey)
	return &SendGrid{client, fromName, fromMail}
}

func (s *SendGrid) SendMail(to, subject, body string) (*Response, error) {
	message := mail.NewSingleEmail(
		mail.NewEmail(s.fromName, s.fromMail),
		subject,
		mail.NewEmail("", to),
		body,
		body,
	)
	res, err := s.client.Send(message)
	if err != nil {
		log.Err(err).Msg("failed to send mail")
		return nil, err
	} else if res.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("got status code from sendgrid: %d", res.StatusCode)
		log.Err(err).Msg("error when sending email")
		return nil, err
	}
	return &Response{StatusCode: res.StatusCode, Body: res.Body}, nil
}
