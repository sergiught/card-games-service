package deck

import (
	"context"
	"database/sql"
	"encoding/json"
)

// RepositoryOperator is an interface with allowed deck repository methods.
type RepositoryOperator interface {
	CreateDeck(ctx context.Context, deck *Deck) error
	OpenDeck(ctx context.Context, deckID string) (*Deck, error)
}

// Repository allows for interacting with
// the card decks stored in the database.
type Repository struct {
	database *sql.DB
}

// NewRepository creates a new Repository
// with a given database connection.
func NewRepository(db *sql.DB) RepositoryOperator {
	return &Repository{database: db}
}

// CreateDeck inserts a new card deck into the database using the deck information provided.
func (r *Repository) CreateDeck(ctx context.Context, deck *Deck) error {
	query := `
		INSERT INTO card_decks
			(deck_id, shuffled, remaining, cards)
		VALUES
			($1, $2, $3, $4)
	`

	_, err := r.database.ExecContext(ctx, query, deck.ID, deck.Shuffled, deck.Remaining, deck.Cards)

	return err
}

// OpenDeck retrieves a card deck from the database using the deck ID.
func (r *Repository) OpenDeck(ctx context.Context, deckID string) (*Deck, error) {
	query := `SELECT deck_id, shuffled, remaining, cards FROM card_decks WHERE deck_id = $1`

	row := r.database.QueryRowContext(ctx, query, deckID)

	var deck Deck
	var cardsJSON string

	if err := row.Scan(&deck.ID, &deck.Shuffled, &deck.Remaining, &cardsJSON); err != nil {
		return nil, err
	}

	var frenchCards []FrenchCard
	if err := json.Unmarshal([]byte(cardsJSON), &frenchCards); err != nil {
		return nil, err
	}

	cards := make([]Card, len(frenchCards))
	for index, card := range frenchCards {
		cards[index] = Card(card)
	}

	deck.Cards = cards

	return &deck, nil
}
