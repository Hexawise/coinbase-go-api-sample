package main

import (
// 	"encoding/json"
	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"reflect"

	"github.com/cucumber/godog"
)

// func (a *apiFeature) resetResponse(*godog.Scenario) {
// 	a.resp = httptest.NewRecorder()
// }

// 	// handle panic
// 	defer func() {
// 		switch t := recover().(type) {
// 		case string:
// 			err = fmt.Errorf(t)
// 		case error:
// 			err = t
// 		}
// 	}()

// 	switch endpoint {
// 	case "/version":
// 		getVersion(a.resp, req)
// 	default:
// 		err = fmt.Errorf("unknown endpoint: %s", endpoint)
// 	}
// 	return
// }

// func (a *apiFeature) theResponseCodeShouldBe(code int) error {
// 	if code != a.resp.Code {
// 		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
// 	}
// 	return nil
// }

// func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
// 	var expected, actual interface{}

// 	// re-encode expected response
// 	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
// 		return
// 	}

// 	// re-encode actual response too
// 	if err = json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
// 		return
// 	}

// 	// the matching may be adapted per different requirements.
// 	if !reflect.DeepEqual(expected, actual) {
// 		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
// 	}
// 	return nil
// }

func iCreateAnOrderRequest(orderType string) error {
	options.Type = orderType

	return nil
}

func theResponseShouldBeSuccessful() error {
	err := CreateOrder(options)
	if err != nil {
		return fmt.Errorf("expected successful response, instead got %v", err.Error())
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

	ctx.Step(`^I create a "(market|limit)" order`, iCreateAnOrderRequest)
	ctx.Step(`^the response should be successful`, theResponseShouldBeSuccessful)
}