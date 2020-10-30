# file: create_order.feature
Feature: create order
  In order to create a Coinbase Order

  Scenario: creating a market order
    When I create a "market" order
    Then the response should be successful