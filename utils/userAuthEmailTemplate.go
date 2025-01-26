package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

type EmailDataType struct {
	SiteName         string
	Username         string
	VerificationLink string
}

type WelcomeEmailDataType struct {
	SiteName string
	Username string
}

// WILL HAVE TO CHANGE THE TEMPLATE
func GenerateEmailContentForVerificationEmailAddress(data EmailDataType) string {
	const emailTemplate = `
   <!DOCTYPE html>
   <html lang="en">
   <head>
   	<meta charset="UTF-8">
   	<meta name="viewport" content="width=device-width, initial-scale=1.0">
   	<title>Email Verification</title>
   	<style>
   		body {
   			font-family: Arial, sans-serif;
   			line-height: 1.6;
   			color: #1a1717;
   			max-width: 600px;
   			margin: 0 auto;
   			padding: 20px;
   		}
   		.container {
   			background-color: #f9f9f9;
   			border: 1px solid #ddd;
   			border-radius: 5px;
   			padding: 20px;
   		}
   	        .button {
                            display: inline-block;
                            padding: 10px 20px;
                            background-color: #115aa8;
                            color: #ffffff; 
                            text-decoration: 
                            border-radius: 5px;
                            text-align: center;
                            text-transform: uppercase;
                            font-weight: bold;
                            font-size: 16px; 
   			 cursor: pointer;
   			 border: none;
   			 rounded: 5px;
   			 margin-top: 10px;
   			 border-radius: 5px;
   
                        }
   
   	</style>
   </head>
   <body>
   	<div class="container">
   		<h2>Welcome to {{.SiteName}}!</h2>
   		<p>Hello {{.Username}},</p>
   		<p>Thank you for registering with us. To complete your registration, please verify your email address by clicking the button below:</p>
   		<p>
   			<a href="{{.VerificationLink}}" class="button">Verify Email</a>
   		</p>
   		<p>Or copy and paste the following link into your browser:</p>
   		<p>{{.VerificationLink}}</p>
   		<p>If you didn't create an account with us, please ignore this email.</p>
   		<p>Best regards,<br>The {{.SiteName}} Team</p>
   	</div>
   </body>
   </html> `

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		log.Printf("error parsing email template: %v", err)
		return ""
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		fmt.Printf("error executing template: %v", err)
		return ""
	}

	return body.String()
}

// will have to change the template
func GenerateEmailContentForResetPassword(data EmailDataType) string {
	const emailTemplate = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Password Reset</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				line-height: 1.6;
				color: #1a1717;
				max-width: 600px;
				margin: 0 auto;
				padding: 20px;
			}
			.container {
				background-color: #f9f9f9;
				border: 1px solid #ddd;
				border-radius: 5px;
				padding: 20px;
			}
			.button {
				display: inline-block;
				padding: 10px 20px;
				background-color: #115aa8;
				color: #ffffff;
				text-decoration: none;
				border-radius: 5px;
				text-align: center;
				text-transform: uppercase;
				font-weight: bold;
				font-size: 16px;
				cursor: pointer;
				border: none;
				rounded: 5px;
				margin-top: 10px;
				border-radius: 5px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Password Reset</h2>
			<p>Hello {{.Username}},</p>
			<p>We received a request to reset your password. If you did not make this request, please ignore this email.</p>
			<p>To reset your password, click the button below:</p>
			<p>
				<a href="{{.VerificationLink}}" class="button">Reset Password</a>
			</p>
			<p>Or copy and paste the following link into your browser:</p>
			<p>{{.VerificationLink}}</p>
			<p>If you did not create an account with us, please ignore this email.</p>
			<p>Best regards,<br>The {{.SiteName}} Team</p>
		</div>
	</body>
	</html>`

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		log.Printf("error parsing email template: %v", err)
		return ""
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		fmt.Printf("error executing template: %v", err)
		return ""
	}

	return body.String()

}

func WelcomeEmailTemplate(data WelcomeEmailDataType) string {
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Welcome to %s</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f9f9f9;
					color: #333;
					line-height: 1.6;
					margin: 0;
					padding: 20px;
				}
				.container {
					max-width: 600px;
					margin: 0 auto;
					background: #fff;
					padding: 20px;
					border-radius: 8px;
					box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
				}
				h1 {
					color: #007BFF;
				}
				p {
					margin: 0 0 15px;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Welcome to %s, %s!</h1>
				<p>We're excited to have you on board. Feel free to explore and make the most of your experience with us.</p>
				<p>If you have any questions, we're here to help.</p>
				<p>Best regards,</p>
				<p>The %s Team</p>
			</div>
		</body>
		</html>
	`, data.SiteName, data.SiteName, data.Username, data.SiteName)
}
