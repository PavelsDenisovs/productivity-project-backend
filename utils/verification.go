package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/resend/resend-go/v2"
)

var randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateVerificationCode() string {
	return fmt.Sprintf("%06d", randomGenerator.Intn(1000000))
}

func SendVerificationEmail(email, code string) error {
	apiKey := os.Getenv("RESEND_API_KEY")
	sender := os.Getenv("SENDER_EMAIL")

	if apiKey == "" || sender == "" {
		return fmt.Errorf("missing RESEND_API_KEY or SENDER_EMAIL")
	}

	client := resend.NewClient(apiKey)

	subject := "Verify Your Email"
	html := fmt.Sprintf("<p>Your verification code is: <strong>%s</strong></p><p>This code will expire in 15 minutes.</p>", code)

	params := &resend.SendEmailRequest{
		From:    sender,
		To:      []string{email},
		Subject: subject,
		Html:    html,
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
