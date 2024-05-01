package scenario

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages/go/v21"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/internal/config"
	"github.com/sergiught/card-games-service/internal/deck"
)

// DeckContext for the various feature test scenarios.
type DeckContext struct {
	config *config.Specification

	database *sql.DB

	response             map[string]interface{}
	checkCardsInResponse bool
}

// NewDeckContext returns a new DeckContext.
func NewDeckContext(cfg *config.Specification, db *sql.DB) *DeckContext {
	return &DeckContext{
		config:   cfg,
		database: db,
	}
}

func (deckCtx *DeckContext) sendRequest(ctx context.Context, method, uri string, body []byte) *http.Response {
	request, err := http.NewRequestWithContext(ctx, method, deckCtx.buildURL(uri), bytes.NewBuffer(body))
	require.NoError(godog.T(ctx), err)

	response, err := http.DefaultClient.Do(request)
	require.NoError(godog.T(ctx), err)

	return response
}

func (deckCtx *DeckContext) buildURL(uri string) string {
	return fmt.Sprintf("http://%s/%s", strings.Trim(deckCtx.config.Server.Address, "/"), strings.Trim(uri, "/"))
}

func (deckCtx *DeckContext) parseCardsTable(ctx context.Context, table *godog.Table) []deck.FrenchCard {
	var cards []deck.FrenchCard

	headers := table.Rows[0].Cells
	for _, row := range table.Rows[1:] {
		card := deck.FrenchCard{
			Value: row.Cells[getTableColumnIndex(ctx, headers, "value")].Value,
			Suit:  row.Cells[getTableColumnIndex(ctx, headers, "suit")].Value,
			Code:  row.Cells[getTableColumnIndex(ctx, headers, "code")].Value,
		}
		cards = append(cards, card)
	}

	return cards
}

func getTableColumnIndex(ctx context.Context, cells []*messages.PickleTableCell, header string) int {
	for idx, cell := range cells {
		if cell.Value == header {
			return idx
		}
	}

	godog.T(ctx).Fatalf("header not found in table: %q", header)

	return 0
}
