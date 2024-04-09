package services

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	FromAddr         = "loggr@dahlton.org"
	FromName         = "Loggr Support"
	ForgotPWTemplate = "d-aa25aab5cda3468db8505bb30209cd47"
)

func NewMailService(sendGridKey string) *MailService {
	return &MailService{
		sendGridKey: sendGridKey,
	}
}

type MailService struct {
	sendGridKey string
}

func (ms *MailService) SendResetPassword(addr, resetLink string) (string, error) {
	request := sendgrid.GetRequest(ms.sendGridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	Body := resetPasswordTemplate(addr, resetLink)
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		return response.Body, err
	}
	return response.Body, nil
}

func resetPasswordTemplate(addr, resetLink string) []byte {
	m := mail.NewV3Mail()

	e := mail.NewEmail(FromName, FromAddr)
	m.SetFrom(e)

	m.SetTemplateID(ForgotPWTemplate)

	p := mail.NewPersonalization()
	tos := []*mail.Email{
		mail.NewEmail("World-Class Gardener", addr),
	}
	p.AddTos(tos...)

	p.SetDynamicTemplateData("url", resetLink)

	m.AddPersonalizations(p)
	return mail.GetRequestBody(m)
}
