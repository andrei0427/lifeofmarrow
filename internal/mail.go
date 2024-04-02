package internal

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string, ccDaniel bool) error {
	smtpUser, ok := os.LookupEnv("SMTP_USER")
	if !ok {
		return errors.New("smtp user not found")
	}

	smtpPass, ok := os.LookupEnv("SMTP_PASS")
	if !ok {
		return errors.New("smtp pass not found")
	}

	smtpHost, ok := os.LookupEnv("SMTP_HOST")
	if !ok {
		return errors.New("smtp host not found")
	}

	smtpPort, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		return errors.New("smtp port not found")
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	headers := make(map[string]string)
	headers["From"] = smtpUser
	headers["To"] = to
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "\r\n\r\n" + body
	msg := []byte(message)

	host := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", host, tlsConfig)
	if err != nil {
		return fmt.Errorf("failure sending smtp email: %s", err)
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failure sending smtp email: %s", err)
	}
	defer client.Quit()

	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("failure sending smtp email: %s", err)
	}

	recipients := []string{to}
	if ccDaniel {
		dan, ok := os.LookupEnv("CONTACT_FORM_TO")
		if !ok {
			return errors.New("smtp to address not found")
		}

		recipients = append(recipients, dan)
	}

	for _, recipient := range recipients {
		if len(recipient) == 0 {
			continue
		}

		err = client.Mail(smtpUser)
		if err != nil {
			return fmt.Errorf("failure sending smtp email: %s", err)
		}

		err = client.Rcpt(recipient)
		if err != nil {
			return fmt.Errorf("failure sending smtp email: %s", err)
		}

		ew, err := client.Data()
		if err != nil {
			return fmt.Errorf("failure sending smtp email: %s", err)
		}

		_, err = ew.Write(msg)
		if err != nil {
			return fmt.Errorf("failure sending smtp email: %s", err)
		}

		err = ew.Close()
		if err != nil {
			return fmt.Errorf("failure sending smtp email: %s", err)
		}
	}

	return nil
}
