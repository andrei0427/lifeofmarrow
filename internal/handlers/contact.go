package handlers

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/partial"
)

func HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		pages.Contact().Render(r.Context(), w)
		return
	}

	// Contact form submission
	form := partial.ContactFields{
		FirstName: r.FormValue("first-name"),
		LastName:  r.FormValue("last-name"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone-number"),
		Message:   r.FormValue("message"),
	}

	errors := make(map[string]string)

	if len(form.FirstName) == 0 {
		errors["first-name"] = "First name is required"
	}

	if len(form.LastName) == 0 {
		errors["last-name"] = "Last name is required"
	}

	if len(form.Email) == 0 {
		errors["email"] = "Email address is required"
	}
	_, err := mail.ParseAddress(form.Email)
	if err != nil {
		errors["email"] = "Invalid email address provided"
	}

	if len(form.Phone) == 0 {
		errors["phone"] = "Phone number is required"
	}

	if len(form.Message) == 0 {
		errors["message"] = "Message is required"
	}

	if len(errors) > 0 {
		partial.ContactForm(form, &errors, nil).Render(r.Context(), w)
		return
	}

	success := false
	smtpUser, ok := os.LookupEnv("SMTP_USER")
	if !ok {
		fmt.Println("Smtp user not found")
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	smtpPass, ok := os.LookupEnv("SMTP_PASS")
	if !ok {
		fmt.Println("Smtp pass not found")
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	smtpHost, ok := os.LookupEnv("SMTP_HOST")
	if !ok {
		fmt.Println("Smtp host not found")
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	smtpPort, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		fmt.Println("Smtp port not found")
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	to, ok := os.LookupEnv("CONTACT_FORM_TO")
	if !ok {
		fmt.Println("Smtp to address not found")
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	bodyParts := make(map[string]string)
	bodyParts["FirstName"] = form.FirstName
	bodyParts["LastName"] = form.LastName
	bodyParts["Email"] = form.Email
	bodyParts["Phone"] = form.Phone
	bodyParts["Message"] = form.Message

	body := ""
	for k, v := range bodyParts {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	headers := make(map[string]string)
	headers["From"] = smtpUser
	headers["To"] = to
	headers["Subject"] = "New Contact Form Submission"

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
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}
	defer client.Quit()

	if err := client.Auth(auth); err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	err = client.Mail(smtpUser)
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	err = client.Rcpt(to)
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	ew, err := client.Data()
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	_, err = ew.Write(msg)
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	err = ew.Close()
	if err != nil {
		fmt.Printf("Failure sending smtp email: %s\n", err)
		partial.ContactForm(form, &errors, &success).Render(r.Context(), w)
		return
	}

	success = true

	partial.ContactForm(form, &errors, &success).Render(r.Context(), w)

}
