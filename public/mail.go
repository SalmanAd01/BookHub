package public

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>Click The Link To Login</p>

`

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func SendMail(to string, subject string, body string) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = GoDotEnvVariable("EMAIL")
	server.Password = GoDotEnvVariable("PASSWORD")
	server.Encryption = mail.EncryptionTLS
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
func main() {
	fmt.Println("Hello")
}
