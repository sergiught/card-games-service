package deck

import (
	"context"
	"database/sql"
)

// RepositoryOperator is an interface with allowed deck repository methods.
type RepositoryOperator interface {
	CreateDeck(ctx context.Context, deck *Deck) error
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
