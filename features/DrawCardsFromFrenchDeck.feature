Feature: Draw cards from a deck of French cards
  As a card dealer
  I want to draw cards from a deck of French cards
  So that I can host a round of Poker

  Scenario: Draw cards from a deck of French cards
    Given I create a "standard" and "sorted" deck of French cards
    And I should see that a deck was created with "52" "sorted" cards
    When I draw "5" cards from the deck
    Then I should see that I received "5" cards from the deck
