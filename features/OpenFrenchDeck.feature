Feature: Open a deck of French cards
  As a card dealer
  I want to open a deck of French cards
  So that I can host a round of Poker

  Scenario: Open a standard deck of French cards
    Given I create a "standard" deck of French cards
    When I open the deck
    Then I should receive "52" cards
