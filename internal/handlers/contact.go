package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"os"

	"github.com/andrei0427/lifeofmarrow/internal"
	"github.com/andrei0427/lifeofmarrow/view/pages"
	"github.com/andrei0427/lifeofmarrow/view/partial"
	"github.com/stripe/stripe-go/v76"
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

	err = internal.SendEmail("", "New Contact Form Submission", body, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending email: %v\n", err)
		partial.ContactForm(form, &errors, stripe.Bool(false)).Render(r.Context(), w)
		return
	}

	partial.ContactForm(form, &errors, stripe.Bool(true)).Render(r.Context(), w)

}
