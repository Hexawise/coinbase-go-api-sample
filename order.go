package main

import (
	"errors"
	"strconv"

	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func CreateOrder(options Options) error {
	var order coinbasepro.Order

	if(options.Type == "limit") {
		order = CreateLimitOrder(options)
	} else {
		order = CreateMarketOrder(options)
	}
	savedOrder, err := client.CreateOrder(&order)

	if err != nil {
		return err
	}

	if savedOrder.ID == "" {
		return errors.New("No create id found")
	}

	return nil
}

func CreateLimitOrder(options Options) coinbasepro.Order {
	postOnly, _ := strconv.ParseBool(options.PostOnly)

	return coinbasepro.Order{
		ClientOID:		options.ClientOrderID,
		Type:					options.Type,
		Side:					options.Side,
		ProductID:		options.ProductID,
		Stp:					options.STP,
		Stop:					options.Stop,
		StopPrice:		options.StopPrice,
		Price:				"0.10",
		Size:					"1.00",
		TimeInForce:	options.TimeInForce,
		CancelAfter:	options.CancelAfter,
		PostOnly:			postOnly,
	}
}

func CreateMarketOrder(options Options) coinbasepro.Order {
	return coinbasepro.Order{
		ClientOID:		options.ClientOrderID,
		Type:					options.Type,
		Side:					options.Side,
		ProductID:		options.ProductID,
		Stp:					options.STP,
		Stop:					options.Stop,
		StopPrice:		options.StopPrice,
		Price:				"0.10",
		Size:					"0.01",
		Funds:				options.Funds,
	}
}