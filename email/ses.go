package email

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SES struct {
}

func InitSES() *SES {
	s := &SES{}
	s.verifyTemplates()

	return s
}

func (s *SES) verifyTemplates() {
	// TODO: check whether ses templates are uploaded
}

func (s *SES) prepareError(err error) error {
	if aerr, ok := err.(awserr.Error); ok {
		// handle as an aws error
		return errors.New(fmt.Sprint("ses: ", aerr.Code(), aerr.Error()))
	}

	return err
}

func (s *SES) Send(mailTo string, template string, data string, mailFrom string) (bool, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an SES client
	svc := ses.New(sess)

	// Define the email parameters
	params := &ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(mailTo),
			},
		},
		Source:   aws.String(mailFrom),
		Template: aws.String(template),
		TemplateData: aws.String(`{
			"subject": "Hello from AWS SES",
			"body": "<p>Hello world!</p>"
		}`),
	}

	// Send the templated email
	_, err := svc.SendTemplatedEmail(params)

	if err != nil {
		return false, s.prepareError(err)
	}

	return true, nil
}
