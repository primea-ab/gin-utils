package elk46

import (
	"github.com/rs/zerolog/log"
)

type SmsLog struct {
	To      string
	Message string
}

type Dummy struct {
	log []*SmsLog
}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (d *Dummy) SendSms(to, message string) (*Response, error) {
	log.Debug().Msgf("Pretending to send sms %s, %s", to, message)
	d.log = append([]*SmsLog{{to, message}}, d.log...)
	return &Response{
		Id:      "never_sent",
		To:      to,
		Message: message,
	}, nil
}

func (d *Dummy) LastMessage() *SmsLog {
	if len(d.log) > 0 {
		return d.log[0]
	}
	return nil
}
