package services

import (
	"fmt"
	"log"
	"miniProjectInterface/constant"
	"miniProjectInterface/entities"
	"miniProjectInterface/externalservices"
)

type OrderServices struct {
	NotifierService externalservices.NotifierService
}

func NewOrderService(notificationType constant.NotificationType) OrderServices {
	return OrderServices{NotifierService: externalservices.NewNotifier(notificationType)}
}

func (o OrderServices) CreatedOrder(order *entities.Order) bool {
	fmt.Printf("order created ==> %+v\n", order)
	// send sms or email
	if o.NotifierService != nil {
		o.NotifierService.Notify(order.UserId, "order created")
	} else {
		log.Fatalln("notifier Service is nil")
	}
	return true
}
