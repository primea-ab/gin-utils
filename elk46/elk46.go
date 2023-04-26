package elk46

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

// https://46elks.se/docs/send-sms
type Elk46 interface {
	SendSms(to, message string) (*Response, error)
}

type Sender struct {
	from          string
	username      string
	password      string
	whendelivered string
	dryrun        string
}

// from  The sender of the SMS as seen by the recipient.
// Either a text sender ID or a phone number in E.164 format if you want to be able to receive replies.
// username for accessing the api
// password for accessing the api
// callbackurl to handle delivery reports. Can be blank to skip callbacks.
// dryrun yes or no
func NewSender(from, username, password, callbackurl, dryrun string) Elk46 {
	if from == "" || username == "" {
		log.Fatal().Msg("missing 46 elk credentials. Cannot create sender...")
	}
	if dryrun != "yes" && dryrun != "no" {
		log.Fatal().Msg("Elk 46 dryrun can only be yes or no")
	}
	return &Sender{
		from, username, password, callbackurl, dryrun,
	}
}

func (s *Sender) SendSms(to, message string) (*Response, error) {
	data := url.Values{
		"from":    {s.from},
		"to":      {to},
		"message": {message},
		"dryrun":  {s.dryrun},
	}
	if s.whendelivered != "" {
		data.Add("whendelivered", s.whendelivered)
	}

	req, _ := http.NewRequest(http.MethodPost, "https://api.46elks.com/a1/SMS", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.SetBasicAuth(s.username, s.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Err(err).Msg("error when sending sms")
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Err(err).Msg("failed to read elk response body")
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Err(err).Msg("failed to Unmarshal elk response")
		return nil, err
	}
	return &response, nil
}
