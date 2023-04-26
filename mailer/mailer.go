package mailer

type Mailer interface {
	SendMail(to, subject, body string) (*Response, error)
}

type Response struct {
	StatusCode int
	Body       string
}
