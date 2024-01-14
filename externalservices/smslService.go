package externalservices

import "fmt"

type SmsService struct {
}

func (s SmsService) Notify(receiver string, message string) {
	fmt.Printf("send sms : { reciever : %s , message : %s }\n", receiver, message)
}
