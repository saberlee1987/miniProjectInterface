package entities

import "miniProjectInterface/constant"

type Order struct {
	Id               int
	UserFullName     string
	UserEmail        string
	UserPhone        string
	UserId           string
	NotificationType constant.NotificationType
	Price            float64
	Status           bool
}
