package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

func setOrderType(_, orderType string) error {
	options.Type = orderType

	return nil
}

func setPriceAndProductID(_, price, baseCurrency, quoteCurrency string) error {
	productPrice := ""

	switch price {
	case "less than the current price":
		productPrice = "0.9"
	case "the current price":
		productPrice = "1.0"
	case "more than the current price":
		productPrice = "1.1"
	}

	options.Price = productPrice
	options.ProductID = baseCurrency + "-" + quoteCurrency

	return nil
}

func setSideSizeAndCurrency(side, size, currency string) error {
	purchaseSize := ""

	switch size {
	case "the smallest amount":
		purchaseSize = "0.1"
	case "the maximum amount":
		purchaseSize = "2"
	default:
		purchaseSize = size
	}

	options.Side = side
	options.Size = purchaseSize

	return nil
}

func theResponseShouldBeSuccessful() error {
	err := CreateOrder(options)
	if err != nil {
		return fmt.Errorf("expected successful response, instead got %v - %+v", err.Error(), options)
	}

	return nil
}

func theResponseShouldNotBeSuccessful() error {
	err := CreateOrder(options)
	if err == nil {
		return fmt.Errorf("expected unsuccessful response, instead got success with %+v", options)
	}

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { options = Options{} })
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		options = Options{} // reset the options before each scenario
	})

	ctx.Step(`^I am (buying|selling) a "([^"]*)" order$`, setOrderType)
	ctx.Step(`^I want to "([^"]*)" "([^"]*)" "([^"]*)"$`, setSideSizeAndCurrency)
	ctx.Step(`^I am willing to (pay|receive) "([^"]*)" of "([^"]*)"\/"([^"]*)"$`, setPriceAndProductID)
	ctx.Step(`^the order should be placed`, theResponseShouldBeSuccessful)
	ctx.Step(`^the order should not be placed`, theResponseShouldNotBeSuccessful)
}