package public

import (
	"fmt"
	"log"

	mail "github.com/xhit/go-simple-mail"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func SendMail(to string, subject string, body string) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "salmanad5s3@gmail.com"
	server.Password = "Sa9860679879@gggggg"
	server.Encryption = mail.EncryptionTLS
	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}
	email := mail.NewMSG()
	email.AddTo(to)
	email.SetSubject(subject)

	email.SetBody(mail.TextHTML, htmlBody)
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
func main() {
	fmt.Println("Hello")
}
