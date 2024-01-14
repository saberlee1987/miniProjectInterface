package main

import (
	"fmt"
	"miniProjectInterface/constant"
	"miniProjectInterface/entities"
	"miniProjectInterface/services"
)

func main() {

	order1 := &entities.Order{
		Id:               1000,
		UserId:           "1500",
		UserFullName:     "saber66",
		UserEmail:        "saberazizi66@yahoo.com",
		UserPhone:        "09365627895",
		Price:            15600000,
		Status:           true,
		NotificationType: constant.Email,
	}

	order2 := &entities.Order{
		Id:               2000,
		UserId:           "2500",
		UserFullName:     "bruce",
		UserEmail:        "bruce40@yahoo.com",
		UserPhone:        "09351298857",
		Price:            985000000,
		Status:           true,
		NotificationType: constant.Sms,
	}

	order3 := &entities.Order{
		Id:           2000,
		UserId:       "2500",
		UserFullName: "bruce",
		UserEmail:    "bruce40@yahoo.com",
		UserPhone:    "09351298857",
		Price:        985000000,
		Status:       true,
	}
	orderService1 := services.NewOrderService(order1.NotificationType)
	orderService2 := services.NewOrderService(order2.NotificationType)
	orderService3 := services.NewOrderService(order3.NotificationType)
	result := orderService1.CreatedOrder(order1)
	fmt.Println("result create order ===> ", result)
	result = orderService2.CreatedOrder(order2)
	fmt.Println("result create order ===> ", result)

	result = orderService3.CreatedOrder(order3)
	fmt.Println("result create order ===> ", result)
}
