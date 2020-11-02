# "Coinbase: Trading with Limit Orders", Revision 60
Feature: Buying using limit orders on CoinBase

	Background:
		Given I am buying a "limit" order

	Scenario Outline: Buying 0
		Given I want to "buy" "0" "<base_currency>"
		And I am willing to pay "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should not be placed
		Examples:
			| side | size | base_currency | price                       | quote_currency |
			| buy  | 0    | ETH           | more than the current price | BTC            |
			| buy  | 0    | BTC           | more than the current price | USD            |
			| buy  | 0    | BAT           | more than the current price | USDC           |
			| buy  | 0    | BTC           | the current price           | USD            |
			| buy  | 0    | BTC           | less than the current price | EUR            |
			| buy  | 0    | LINK          | more than the current price | USDC           |

	Scenario Outline: Buying the smallest amount
		Given I want to "buy" "the smallest amount" "<base_currency>"
		And I am willing to pay "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size                | base_currency | price                       | quote_currency |
			| buy  | the smallest amount | BTC           | the current price           | EUR            |
			| buy  | the smallest amount | BAT           | less than the current price | USDC           |

	Scenario Outline: Buying 1
		Given I want to "buy" "1" "<base_currency>"
		And I am willing to pay "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size | base_currency | price                       | quote_currency |
			| buy  | 1    | BTC           | less than the current price | USD            |
			| buy  | 1    | BTC           | less than the current price | GBP            |
			| buy  | 1    | ETH           | less than the current price | BTC            |
			| buy  | 1    | BTC           | more than the current price | USD            |
			| buy  | 1    | LINK          | less than the current price | USDC           |

	Scenario Outline: Buying the maximum amount
		Given I want to "buy" "the maximum amount" "<base_currency>"
		And I am willing to pay "<price>" of "<base_currency>"/"<quote_currency>"
		Then the order should be placed
		Examples:
			| side | size               | base_currency | price                       | quote_currency |
			| buy  | the maximum amount | ETH           | the current price           | BTC            |
			| buy  | the maximum amount | LINK          | less than the current price | USDC           |
			| buy  | the maximum amount | BTC           | the current price           | GBP            |
			| buy  | the maximum amount | BTC           | more than the current price | USD            |
			| buy  | the maximum amount | BAT           | the current price           | USDC           |

