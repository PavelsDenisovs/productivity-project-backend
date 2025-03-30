package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", randomGenerator.Intn(1000000))
}

func SendVerificationEmail(email, code string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	senderEmail := os.Getenv("SENDER_EMAIL")

	if smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" || senderEmail == "" {
		return fmt.Errorf("missing SMTP credentials in .env file")
	}

	subject := "Verify Your Email"
	body := fmt.Sprintf("Your verification code is: %s\n\nThis code will expire in 15 minutes.", code)

	message := fmt.Sprintf("From: %s\r\n", senderEmail) +
		fmt.Sprintf("To: %s\r\n", email) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"\r\n" +
		body

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{email}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
