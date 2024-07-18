package main

import (
	"log"
	"time"

	"github.com/rajath002/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	// create and trigger un Named Go-Routine function
	go func() {
		// for loop with no conditions
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}() // trigger while creating
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()

	// SMTP server configuration for MailHog
	server.Host = "localhost"
	// 1025 mailhog port to handle requests, 8025 is the UI/client port
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.Encryption = mail.EncryptionNone

	// Initialize the client
	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	// Create a new email
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	// Send the email
	err = email.Send(client)
	if err != nil {
		log.Println("error while sending mail.Send : ", err)
	} else {
		log.Println("Email Sent!")
	}
}
