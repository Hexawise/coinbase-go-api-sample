package main

import (
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func CreateOrder(options Options) {
	println(options.Type)
	println("Order Created")
}

func ListOrders(options Options) {
	var orders []coinbasepro.Order
	cursor := client.ListOrders()

	println("Orders:")

	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			println("Error:", err.Error())
			return
		}
	
		for _, o := range orders {
			println(o.ID)
		}
	}
}