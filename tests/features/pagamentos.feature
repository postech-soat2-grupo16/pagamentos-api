Feature: API Pagamentos

  Scenario: Receive payment callback from MercadoPago
    Given a MercadoPago payment callback
    Then the payment status should be updated

  Scenario: Get payment by ID
    Given a payment ID
    When the user requests the payment by ID
    Then the API should respond with the payment details

  Scenario: Health check
    When the health endpoint is accessed
    Then the API should respond with "OK"