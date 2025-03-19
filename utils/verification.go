package utils

import (
	"fmt"
	"math/rand"
	"net/smpt"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", randomGenerator.Intn(1000000))
}

func SendVerificationEmail(email, code string) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env: %w", err)
	}


	smtpHost := os.Getenv("SMPT_HOST")
	smtpPort := os.Getenv("SMPT_PORT")
	smtpUsername := os.Getenv("SMPT_USERNAME")
	smtpPassword := os.Getenv("SMPT_PASSWORD")
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