package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
)

func iCreateAnOrderRequest(orderType string) error {
	options.Type = orderType

	return nil
}

func theResponseShouldBeSuccessful() error {
	err := CreateOrder(options)
	if err != nil {
		return fmt.Errorf("expected successful response, instead got %v\n%+v\n", err.Error(), options)
	}

	return nil
}

func iSetTheOptionTo(option, value string) error {
	v := reflect.ValueOf(&options).Elem().FieldByName(option)

	if v.IsValid() {
		v.SetString(value)

		return nil
	} else {
		return fmt.Errorf("expected to set option \"%v\" to \"%v\", but \"%v\" is not a valid option", option, value, option)
	}
}

type DepositCrypto struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id"`
}

type DepositCoinbase struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	ID                string `json:"id"`
}

type PaymentMethod struct {
	PaymentMethodID		string `json:"id"`
	Type							string `json:"profile_id"`
	Currency					string `json:"currency"`
}

func createDeposit(currency, amount string) {
	newDepositCrypto := DepositCrypto{
		Amount: 						amount,
		Currency: 					currency,
		CoinbaseAccountID:	os.Getenv("COINBASE_ACCOUNT_ID"),
	}

	client = NewClient()

	var savedDeposit DepositCoinbase
	_, err := client.Request("POST", "/deposits/coinbase-account", newDepositCrypto, &savedDeposit)
	
	if err != nil {
		println("ERR", err.Error())
		fmt.Errorf("could not deposit: %v", err.Error())
	} else {
		println("SUCC")
		fmt.Printf("%+v", savedDeposit)
	}
}

func createWithdrawal(currency, amount string) {
	newDepositCrypto := DepositCrypto{
		Amount: 						amount,
		Currency: 					currency,
		CoinbaseAccountID:	os.Getenv("COINBASE_ACCOUNT_ID"),
	}

	client = NewClient()

	var savedDeposit DepositCoinbase
	_, err := client.Request("POST", "/withdrawals/coinbase-account", newDepositCrypto, &savedDeposit)
	
	if err != nil {
		println("ERR", err.Error())
		fmt.Errorf("could not deposit: %v", err.Error())
	} else {
		println("SUCC")
		fmt.Printf("%+v", savedDeposit)
	}
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		options = Options{}

		// var paymentMethods []PaymentMethod
		// _, err := client.Request("GET", "/accounts", nil, &paymentMethods)

		// if err != nil {
		// 	println("ERR", err.Error())
		// 	fmt.Errorf("could not deposit: %v", err.Error())
		// } else {
		// 	println("SUCC")
		// 	fmt.Printf("%+v", paymentMethods)
		// }

		// createDeposit("BTC", "50.0")
		createWithdrawal("BTC", "40.0")
		// createDeposit("ETH", "500.0")
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		options = Options{} // reset the options before each scenario
	})

	ctx.Step(`^I want to create a "(market|limit)" order`, iCreateAnOrderRequest)
	ctx.Step(`^I set the "([^"]*)" option to "([^"]*)`, iSetTheOptionTo)
	ctx.Step(`^the response should be successful`, theResponseShouldBeSuccessful)
}