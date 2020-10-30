package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	// General
	ClientOrderID string `long:"client_oid" description:"Order ID selected by you to identify your order"`
	Type string `long:"type" description:"limit or market (default is limit)" default:"limit" choice:"limit" choice"market"`
	Side string `long:"side" description:"buy or sell" required:"true"`
	ProductID string `long:"product_id" description:"A valid product id" required:"True"`
	STP string `long:"stp" description:" Self-trade prevention flag"`
	Stop string `long:"stop" description:"Either loss or entry. Requires stop_price to be defined"`
	StopPrice string `long:"stop_price" description:"Only if stop is defined. Sets trigger price for stop order"`

	// Limit Order
	Price string `long:"price" description:"[Limit] Price per bitcoin"`
	Size string `long:"size" description:"[Limit|Market] Amount of base currency to buy or sell"`
	TimeInForce string `long:"time_in_force" description:"[Limit] [optional] GTC, GTT, IOC, or FOK (default is GTC)" default:"GTC" choice:"GTC" choice:"GTT" choice:"IOC" choice:"FOK"`
	CancelAfter string `long:"cancel_after" description:"[Limit] [optional] min, hour, day, requires time_in_force to be GTT"`
	PostOnly string `long:"post_only" description:"[Limit] [optional] Post only flag", invalid when time_in_force is IOC or FOK`

	// Market Order
	Funds string `long:"funds" description:"[Market] [optional] Desired amount of quote currency to use, can only use size or funds"`
}

var options Options

var client = NewClient()

func main() {
	_, err := flags.ParseArgs(&options, os.Args[1:])

	if err != nil {
		if err.(*flags.Error).Type != flags.ErrHelp {
			panic(err)
		}
	} else {
		ListProducts(options)
		println("----")
		CreateOrder(options)
		println("----")
		ListOrders(options)
	}
}