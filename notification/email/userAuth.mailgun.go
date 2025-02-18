package notification

import (
	"carveo/config"
	"carveo/utils"
	"context"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type mailgunService struct {
	mailGunConfig config.MailgunConfig
}

func NewMailgunService() *mailgunService {
	cfg := config.GetConfig() // Get the globally loaded config
	fmt.Println("Mailgun config:", cfg)
	if cfg == nil {
		fmt.Println("Mailgun config is not initialized")
		return nil
	}

	fmt.Printf("Mailgun Service Initialized with Config: %+v\n", cfg.MailgunConfig)

	return &mailgunService{
		mailGunConfig: cfg.MailgunConfig,
	}
}

func (mg *mailgunService) SendVerificationEmail(to string, token string, username string, verificationBaseUrl string) error {
	mgClient := mailgun.NewMailgun(mg.mailGunConfig.Domain, mg.mailGunConfig.PrivateAPIKey)
	data := &utils.EmailDataType{}

	data = &utils.EmailDataType{
		SiteName:         "Carveo",
		Username:         username,
		VerificationLink: fmt.Sprintf("%s/api/v1/userAuth/verifyEmail?token=%s", verificationBaseUrl, token),
	}

	body := utils.GenerateEmailContentForVerificationEmailAddress(*data)
	subject := "Verify Your Email Address"
	message := mgClient.NewMessage(mg.mailGunConfig.SenderEmail, subject, body, to)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mgClient.Send(ctx, message)
	return err
}

func (mg *mailgunService) SendResetPasswordEmail(to string, token string, username string, verificationBaseUrl string) error {
	mgClient := mailgun.NewMailgun(mg.mailGunConfig.Domain, mg.mailGunConfig.PrivateAPIKey)
	data := &utils.EmailDataType{}

	data = &utils.EmailDataType{
		SiteName:         "Carveo",
		Username:         username,
		VerificationLink: fmt.Sprintf("%s/api/userAuth/validateResetPasswordToken?token=%s", verificationBaseUrl, token),
	}

	body := utils.GenerateEmailContentForResetPassword(*data)
	subject := "Reset Your Password"
	message := mgClient.NewMessage(mg.mailGunConfig.SenderEmail, subject, body, to)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mgClient.Send(ctx, message)
	return err
}

func (mg *mailgunService) SendWelcomeEmail(to string, userName string) error {
	mgClient := mailgun.NewMailgun(mg.mailGunConfig.Domain, mg.mailGunConfig.PrivateAPIKey)

	data := &utils.WelcomeEmailDataType{
		SiteName: "Carveo",
		Username: to,
	}

	body := utils.WelcomeEmailTemplate(*data)
	subject := "Welcome to Carveo"
	message := mgClient.NewMessage(mg.mailGunConfig.SenderEmail, subject, body, to)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mgClient.Send(ctx, message)
	return err
}
