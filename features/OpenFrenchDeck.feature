Feature: Open a deck of French cards
  As a card dealer
  I want to open a deck of French cards
  So that I can host a round of Poker

  Scenario: Open a standard deck of French cards
    Given I create a "standard" and "sorted" deck of French cards
    And I should see that a deck was created with "52" "sorted" cards
    When I open the deck
    Then I should see that the deck has the following "sorted" cards:
      | value | suit    | code |
      | ACE   | SPADES  | AS   |
      | 2     | SPADES  | 2S   |
      | 3     | SPADES  | 3S   |
      | 4     | SPADES  | 4S   |
      | 5     | SPADES  | 5S   |
      | 6     | SPADES  | 6S   |
      | 7     | SPADES  | 7S   |
      | 8     | SPADES  | 8S   |
      | 9     | SPADES  | 9S   |
      | 10    | SPADES  | 10S  |
      | JACK  | SPADES  | JS   |
      | QUEEN | SPADES  | QS   |
      | KING  | SPADES  | KS   |
      | ACE   | DIAMONDS| AD   |
      | 2     | DIAMONDS| 2D   |
      | 3     | DIAMONDS| 3D   |
      | 4     | DIAMONDS| 4D   |
      | 5     | DIAMONDS| 5D   |
      | 6     | DIAMONDS| 6D   |
      | 7     | DIAMONDS| 7D   |
      | 8     | DIAMONDS| 8D   |
      | 9     | DIAMONDS| 9D   |
      | 10    | DIAMONDS| 10D  |
      | JACK  | DIAMONDS| JD   |
      | QUEEN | DIAMONDS| QD   |
      | KING  | DIAMONDS| KD   |
      | ACE   | CLUBS   | AC   |
      | 2     | CLUBS   | 2C   |
      | 3     | CLUBS   | 3C   |
      | 4     | CLUBS   | 4C   |
      | 5     | CLUBS   | 5C   |
      | 6     | CLUBS   | 6C   |
      | 7     | CLUBS   | 7C   |
      | 8     | CLUBS   | 8C   |
      | 9     | CLUBS   | 9C   |
      | 10    | CLUBS   | 10C  |
      | JACK  | CLUBS   | JC   |
      | QUEEN | CLUBS   | QC   |
      | KING  | CLUBS   | KC   |
      | ACE   | HEARTS  | AH   |
      | 2     | HEARTS  | 2H   |
      | 3     | HEARTS  | 3H   |
      | 4     | HEARTS  | 4H   |
      | 5     | HEARTS  | 5H   |
      | 6     | HEARTS  | 6H   |
      | 7     | HEARTS  | 7H   |
      | 8     | HEARTS  | 8H   |
      | 9     | HEARTS  | 9H   |
      | 10    | HEARTS  | 10H  |
      | JACK  | HEARTS  | JH   |
      | QUEEN | HEARTS  | QH   |
      | KING  | HEARTS  | KH   |

  Scenario: Throw an error if we open a deck that does not exist
    When I open the deck
    Then I should see an error saying that the deck does not exist
