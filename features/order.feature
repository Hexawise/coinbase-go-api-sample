# "Coinbase: Place a New Order", Revision 39
Feature: I want to create an order on CoinBase

	Scenario Outline: Creating a Limit Order
		Given I want to create a "limit" order
		When I set the "Side" option to "<side>"
		And I set the "ProductID" option to "<product_id>"
		And I set the "Price" option to "<price>"
		And I set the "Size" option to "<size>"
		Then the response should be successful
		Examples:
			| type  | side | product_id    | price               | size |
			| limit | buy  | BTC-USD       | at current price    | none |
			| limit | sell | ETH-BTC       | above current price | many |
			| limit | buy  | BAT-USDC      | above current price | few  |
			| limit | sell | BTC-EUR       | below current price | none |
			| limit | buy  | invalid       | at current price    | many |
			| limit | sell | LINK-USDC     | below current price | few  |
			| limit | buy  | not specified | below current price | many |
			| limit | buy  | ETH-BTC       | above current price | none |
			| limit | sell | BTC-GBP       | above current price | few  |

	Scenario Outline: Creating a Market Order
		Given I want to create a "market" order
		When I set the "Side" option to "<side>"
		And I set the "ProductID" option to "<product_id>"
		Then the response should be successful
		Examples:
			| type   | side | product_id    |
			| market | buy  | ETH-BTC       |
			| market | sell | BTC-GBP       |
			| market | sell | not specified |
			| market | sell | BTC-EUR       |
			| market | sell | invalid       |
			| market | sell | LINK-USDC     |
			| market | buy  | BAT-USDC      |
			| market | sell | BTC-USD       |