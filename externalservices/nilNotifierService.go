package externalservices

import "fmt"

type NilNotifierService struct {
}

func (s NilNotifierService) Notify(receiver string, message string) {
	fmt.Printf("NilNotifierService : { reciever : %s , message : %s }\n", receiver, message)
}
