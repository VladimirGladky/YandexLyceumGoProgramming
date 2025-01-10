package main

import "log"

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

type OrderLogger struct{}

func NewOrderLogger() OrderLogger {
	return OrderLogger{}
}

func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f", order.OrderNumber, order.CustomerName, order.OrderAmount)
}
