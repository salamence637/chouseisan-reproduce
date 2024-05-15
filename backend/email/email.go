package email

import (
	"fmt"
	"net/smtp"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Email struct {
	To      []string
	Subject string
	Body    string
}

func SendEmail(c *gin.Context) {
	// get request body
	var e Email
	if err := c.BindJSON(&e); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	// Set up authentication information.
	from := "q.mashio.p@gmail.com"
	password := "udehhencbegimpal"

	// Set up smtp server information.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Set up receipient information
	to := e.To
	/*
	to := []string{
		"uzenkyu@gmail.com",
	}
	*/

	// Message.
	message := []byte("To: " + strings.Join(to, ",") + "\r\n" + 
		"Subject: " + e.Subject + "\r\n" + 
		"\r\n" + 
		e.Body + "\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Email Sent Successfully")
	return
}
