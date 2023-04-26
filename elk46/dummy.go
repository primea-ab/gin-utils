package elk46

import (
	"github.com/rs/zerolog/log"
)

type Dummy struct{}

func NewDummy() Elk46 {
	return &Dummy{}
}

func (d *Dummy) SendSms(to, message string) (*Response, error) {
	log.Debug().Msgf("Pretending to send sms %s, %s", to, message)
	return &Response{
		Id:      "never_sent",
		To:      to,
		Message: message,
	}, nil
}
