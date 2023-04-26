package mailer

type MailLog struct {
	To      string
	Subject string
	Body    string
}

type Dummy struct {
	log []*MailLog
}

func NewDummy() *Dummy {
	return &Dummy{}
}

func (d *Dummy) SendMail(to, subject, body string) (*Response, error) {
	d.log = append([]*MailLog{{to, subject, body}}, d.log...)
	return &Response{StatusCode: 200, Body: ""}, nil
}

func (d *Dummy) LastMessage() *MailLog {
	if len(d.log) > 0 {
		return d.log[0]
	}
	return nil
}
