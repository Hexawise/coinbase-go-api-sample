# file: create_order.feature
Feature: get version
  In order to know godog version
  As an API user
  I need to be able to request version

  Scenario: not all arguments provided
    When I send "POST" request to "/version"
    Then the response code should be 405
    And the response should match json:
      """
      {
        "error": "Method not allowed"
      }
      """

  Scenario: should get version number
    When I send "GET" request to "/version"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "version": "v0.10.0"
      }
      """