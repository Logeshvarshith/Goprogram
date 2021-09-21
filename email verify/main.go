package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	fmt.Println("Golang email app running")
	email()
}

func email() {
	os.Setenv("FromEmailAddr", "logesh1104@gmail.com")
	os.Setenv("SMTPpwd", "15ec046logesh")
	os.Setenv("ToEmailAddr", "t4snaneclogeshwaran@gmail.com")
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")
	// receiver address
	toEmail := os.Getenv("ToEmailAddr") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Our Golang Email\n"
	body := "our first email!"
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully sent mail to all user in toList")
}
