package externalservices

import "miniProjectInterface/constant"

type NotifierService interface {
	Notify(receiver string, message string)
}

func NewNotifier(notificationType constant.NotificationType) NotifierService {
	switch notificationType {
	case constant.Email:
		return EmailService{}
	case constant.Sms:
		return SmsService{}
	default:
		return NilNotifierService{}
	}
}
