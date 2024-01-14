package externalservices

import "fmt"

type EmailService struct {
}

func (e EmailService) Notify(receiver string, message string) {
	fmt.Printf("send email :  { reciever : %s , message : %s }\n", receiver, message)
}
