package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	EMAIL_PASSWORD = os.Getenv("EMAIL_PASSWORD")
	EMAIL_SERVER   = os.Getenv("EMAIL_SERVER")
	EMAIL_PORT     = os.Getenv("EMAIL_PORT")
)

type Email struct {
	Sender     string   `json:"sender"`
	Pass       string   `json:"pass,omitempty"`
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
}

// Handles all of the sending logic
func (e *Email) Send(server, port string) (err error) {

	// Auth via Sender app password
	auth := smtp.PlainAuth("", e.Sender, e.Pass, server)

	err = smtp.SendMail(
		fmt.Sprintf("%v:%v", server, port),
		auth,
		e.Sender,
		e.Recipients,
		[]byte(fmt.Sprintf("Subject: %s \r\n\r\n%s", e.Subject, e.Body)),
	)

	for _, i := range e.Recipients {
		log.Printf("%s  -->  %s\n", e.Sender, i)
	}

	return
}

func (e *Email) Output() {
	fmt.Printf("Sender  : %s\n", e.Sender)
	fmt.Printf("Subject : %s\n", e.Subject)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func sendEmail(c echo.Context) error {
	var email Email

	if err := c.Bind(&email); err != nil {
		return err
	}

	email.Pass = EMAIL_PASSWORD

	if err := email.Send(EMAIL_SERVER, EMAIL_PORT); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, email)
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/send", sendEmail)

	e.Logger.Info(e.Start("0.0.0.0:9000"))
}
