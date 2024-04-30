package scenario

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/internal/config"
)

// DeckContext for the various feature test scenarios.
type DeckContext struct {
	config *config.Specification

	response map[string]interface{}
}

// NewDeckContext returns a new DeckContext.
func NewDeckContext(cfg *config.Specification) *DeckContext {
	return &DeckContext{
		config: cfg,
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
