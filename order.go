package main

import (
	"errors"

	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func CreateOrder(options Options) (err error) {
	println("Order Created")

	if(false) {
		return errors.New("Some Error")
	} else {
		return
	}
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