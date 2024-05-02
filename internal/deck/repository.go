package deck

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

var (
	// ErrNotEnoughCards is an error that indicates that we are
	// about to draw more cards than available in the deck.
	ErrNotEnoughCards = errors.New("not enough cards in the deck")
)

// RepositoryOperator is an interface with allowed deck repository methods.
type RepositoryOperator interface {
	CreateDeck(ctx context.Context, deck *Deck) error
	OpenDeck(ctx context.Context, deckID string) (*Deck, error)
	DrawCards(ctx context.Context, deckID string, cardsToDraw int) ([]FrenchCard, error)
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

// DrawCards draws N cards from a given card deck in the database.
func (r *Repository) DrawCards(ctx context.Context, deckID string, cardsToDraw int) ([]FrenchCard, error) {
	tx, err := r.database.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	query := `SELECT cards FROM card_decks WHERE deck_id = $1`
	row := tx.QueryRowContext(ctx, query, deckID)

	var cardsJSON string
	if err := row.Scan(&cardsJSON); err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	var frenchCards []FrenchCard
	if err := json.Unmarshal([]byte(cardsJSON), &frenchCards); err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	if len(frenchCards) < cardsToDraw {
		return nil, errors.Join(ErrNotEnoughCards, tx.Rollback())
	}

	drawnCards := frenchCards[:cardsToDraw]
	remainingCards := frenchCards[cardsToDraw:]

	updatedSerializedCards, err := json.Marshal(remainingCards)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	query = `UPDATE card_decks SET cards = $1, remaining = $2 WHERE deck_id = $3`
	_, err = tx.ExecContext(ctx, query, updatedSerializedCards, len(remainingCards), deckID)
	if err != nil {
		return nil, errors.Join(err, tx.Rollback())
	}

	return drawnCards, tx.Commit()
}
