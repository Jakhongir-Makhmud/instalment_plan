package sms

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)


type Sms struct {
	client twilio.RestClient `json:"-"`
}

// initialize Sms service
func NewSmsSerivce(client twilio.RestClient) *Sms {
	return &Sms{
		client: client,
	}
}


// SendSms sends sms only verified numbers, because it uses free Twilio account
// only verified number is mine 
func (s Sms) SendSms(phoneNumber, message string) error {
	msg := &openapi.CreateMessageParams{}
	msg.SetFrom("+19402363343")
	msg.SetTo(phoneNumber)
	msg.SetBody(message)

	_, err := s.client.ApiV2010.CreateMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
