package email

import (
	"carveo/config"
	"carveo/utils"
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type mailgunService struct {
	config.MailgunConfig
}

func NewMailgunService(cfg config.MailgunConfig) *mailgunService {
	return &mailgunService{
		MailgunConfig: cfg,
	}
}

func (mg *mailgunService) SendVerificationEmail(to string, token string, username string, verificationBaseUrl string) error {
	mgClient := mailgun.NewMailgun(mg.Domain, mg.PrivateAPIKey)
	data := &utils.EmailDataType{}

	data = &utils.EmailDataType{
		SiteName:         "Carveo",
		Username:         username,
		VerificationLink: fmt.Sprintf("%s/api/userAuth/verifyEmail?token=%s", verificationBaseUrl, token),
	}

	body := utils.GenerateEmailContentForVerificationEmailAddress(*data)
	subject := "Verify Your Email Address"
	message := mgClient.NewMessage(mg.SenderEmail, subject, body, to)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mgClient.Send(ctx, message)
	return err
}

func (mg *mailgunService) SendResetPasswordEmail(to string, token string, username string, verificationBaseUrl string) error {
	mgClient := mailgun.NewMailgun(mg.Domain, mg.PrivateAPIKey)
	data := &utils.EmailDataType{}

	data = &utils.EmailDataType{
		SiteName:         "Carveo",
		Username:         username,
		VerificationLink: fmt.Sprintf("%s/api/userAuth/resetPassword?token=%s", verificationBaseUrl, token),
	}

	body := utils.GenerateEmailContentForResetPassword(*data)
	subject := "Reset Your Password"
	message := mgClient.NewMessage(mg.SenderEmail, subject, body, to)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mgClient.Send(ctx, message)
	return err

}
