# "Coinbase: Trading with Limit Orders", Revision 62
Feature: Selling using limit orders on CoinBase

	Background: 
		Given I am selling a "limit" order

	Scenario Outline: Selling 0
		Given I want to "sell" "0" "<base_currency>"
		And I am willing to receive "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should not be placed
		Examples:
			| side | size | base_currency | price                       | quote_currency |
			| sell | 0    | BTC           | the current price           | GBP            |
			| sell | 0    | ETH           | the current price           | BTC            |
			| sell | 0    | BAT           | more than the current price | USDC           |
			| sell | 0    | ETH           | less than the current price | BTC            |
			| sell | 0    | ETH           | more than the current price | BTC            |
			| sell | 0    | BTC           | less than the current price | EUR            |

	Scenario Outline: Selling the smallest amount
		Given I want to "sell" "the smallest amount" "<base_currency>"
		And I am willing to receive "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size                | base_currency | price                       | quote_currency |
			| sell | the smallest amount | LINK          | the current price           | USDC           |
			| sell | the smallest amount | BTC           | more than the current price | USD            |
			| sell | the smallest amount | BTC           | more than the current price | GBP            |
			| sell | the smallest amount | BAT           | less than the current price | USDC           |
			| sell | the smallest amount | ETH           | more than the current price | BTC            |
			| sell | the smallest amount | BTC           | the current price           | USD            |
			| sell | the smallest amount | BTC           | less than the current price | USD            |
			| sell | the smallest amount | BTC           | more than the current price | EUR            |
			| sell | the smallest amount | ETH           | less than the current price | BTC            |
			| sell | the smallest amount | BAT           | the current price           | USDC           |

	Scenario Outline: Selling 1
		Given I want to "sell" "1" "<base_currency>"
		And I am willing to receive "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size | base_currency | price                       | quote_currency |
			| sell | 1    | BAT           | less than the current price | USDC           |
			| sell | 1    | BTC           | the current price           | GBP            |
			| sell | 1    | BTC           | less than the current price | EUR            |
			| sell | 1    | LINK          | less than the current price | USDC           |

	Scenario Outline: Selling the maximum amount
		Given I want to "sell" "the maximum amount" "<base_currency>"
		And I am willing to receive "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size               | base_currency | price                       | quote_currency |
			| sell | the maximum amount | BTC           | more than the current price | GBP            |
			| sell | the maximum amount | BTC           | more than the current price | EUR            |
			| sell | the maximum amount | LINK          | more than the current price | USDC           |