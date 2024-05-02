Feature: Draw cards from a deck of French cards
  As a card dealer
  I want to draw cards from a deck of French cards
  So that I can host a round of Poker

  Scenario: Draw cards from a deck of French cards
    Given I create a "standard" and "sorted" deck of French cards
    And I should see that a deck was created with "52" "sorted" cards
    When I draw "5" cards from the deck
    Then I should see that I received "5" cards from the deck

  Scenario: Throw an error if we draw from a deck that does not exist
    When I draw "5" cards from the deck
    Then I should see an error saying that the deck does not exist

  Scenario: Throw an error if there are not enough cards to draw from the deck
    Given I create a "custom" and "sorted" deck of French cards with the following cards:
      | value | suit    | code |
      | ACE   | SPADES  | AS   |
      | 2     | SPADES  | 2S   |
    And I should see that a deck was created with "2" "sorted" cards
    When I draw "5" cards from the deck
    Then I should see an error saying that there are not enough cards in the deck
