package public

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(to string, subject string, body string) {
	from := "salmanad5s3@gmail.com"
	pass := "Sa9860679879@gggggg"
	toList := make([]string, 1)
	toList = append(toList, to)
	host := "smtp.gmail.com"
	port := "587"
	bodymsg := make([]byte, 1)
	bodymsg = append(bodymsg, []byte(body)...)
	auth := smtp.PlainAuth("", from, pass, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, bodymsg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
func main() {
	fmt.Println("Hello")
}
