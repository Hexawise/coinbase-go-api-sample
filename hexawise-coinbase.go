package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	ClientOrderID string `long:"client_oid" description:"Order ID selected by you to identify your order"`
	Type string `long:"type" description:"limit or market (default is limit)" default:"limit"`
	Side string `long:"side" description:"buy or sell" required:"true"`
	ProductID string `long:"product_id" description:"A valid product id" required:"True"`
	STP string `long:"stp" description:" Self-trade prevention flag"`
	Stop string `long:"stop" description:"Either loss or entry. Requires stop_price to be defined"`
	StopPrice string `long:"stop_price" description:"Only if stop is defined. Sets trigger price for stop order"`
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