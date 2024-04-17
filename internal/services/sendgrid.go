package services

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	FromAddr           = "loggr@dahlton.org"
	FromName           = "Loggr Support"
	ForgotPWTemplateId = "d-aa25aab5cda3468db8505bb30209cd47"
	NewUserTemplateId  = "d-d4f1bfd77cd146548d5eb5f0cb669832"
)

func NewMailService(sendGridKey string) *MailService {
	return &MailService{
		sendGridKey: sendGridKey,
	}
}

type MailService struct {
	sendGridKey string
}

type templateFunc func(string, string) []byte

func (ms *MailService) SendEmailWithTemplate(addr, resetLink string, tf templateFunc) (string, error) {
	request := sendgrid.GetRequest(ms.sendGridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	Body := tf(addr, resetLink)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		return response.Body, err
	}
	return response.Body, nil
}

func ResetPasswordTemplate(addr, resetLink string) []byte {
	if addr == "" || resetLink == "" {
		return []byte{}
	}
	m := mail.NewV3Mail()

	e := mail.NewEmail(FromName, FromAddr)
	m.SetFrom(e)

	m.SetTemplateID(ForgotPWTemplateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("World-Class Gardener", addr),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("url", resetLink)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}

func NewUserTemplate(addr, user string) []byte {
	if addr == "" || user == "" {
		return []byte{}
	}
	m := mail.NewV3Mail()

	e := mail.NewEmail(FromName, FromAddr)
	m.SetFrom(e)

	m.SetTemplateID(NewUserTemplateId)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("Our new friend, "+user, addr),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("user", user)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}
