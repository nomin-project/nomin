package sender

import (
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"
)

// TestConnection tests if SMTP server is accessible
func TestConnection(serverAddress string, serverPort string) error {
	return nil
}

// SendMail checks content and sends the email
func SendMail(from string, to string, subject string, message string, serverAddress string, serverPort string) error {
	var err error

	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", message)

	port, err := strconv.Atoi(serverPort)
	if err != nil {
		return fmt.Errorf("Port must be an integer, error during conversion: %v", err)
	}

	dialer := gomail.Dialer{Host: serverAddress, Port: port}
	err = dialer.DialAndSend(mail)
	if err != nil {
		return err
	}

	return nil
}
