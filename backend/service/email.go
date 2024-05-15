package service

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(userName, email, eventID, title string) error {
	from := "chousei208@gmail.com"
	password := os.Getenv("CHOUSEISAN_EMAIL_PASSWORD")

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	body := fmt.Sprintf("Dear %s, \n The poll for your event %s has ended.\n You can check the result here:\n http://localhost:3000/scheduler/view_event/%s", userName, title, eventID)

	// Compose the message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Event Poll Ended!" + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{email}, msg)
	return err
}

func SendEmailAdd(userName, email, eventID, title, due_edit string) error {
	from := "chousei208@gmail.com"
	password := os.Getenv("CHOUSEISAN_EMAIL_PASSWORD")

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	body := fmt.Sprintf("Dear %s, \n You have added your attendance to event %s.\n The deadline for editing your attendance is %s.\n You can check the poll here:\n http://localhost:3000/scheduler/view_event/%s\n\n", userName, title, due_edit, eventID)

	// Compose the message
	msg := []byte("To: " + email + "\r\n" +
		"Subject: Attendance Added!" + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{email}, msg)
	return err
}
