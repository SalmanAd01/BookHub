package controller

import (
	"fmt"
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, There</title>
</head>
<body>
   <p>Click The Link</p>

`

func SendMail(to, subject, body string) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 465
	server.Username = os.Getenv("EMAIL")
	server.Password = os.Getenv("PASSWORD")
	server.Encryption = mail.EncryptionSSL
	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()
	email.AddTo(to)
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, htmlBody+"<p>"+body+"</p></body>")
	err = email.Send(smtpClient)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
