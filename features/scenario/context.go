package scenario

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Context for the various feature test scenarios.
type Context struct {
	testingT *testing.T
	assert   *assert.Assertions
	require  *require.Assertions
}

// NewContext returns a new Context.
func NewContext(t *testing.T) *Context {
	return &Context{
		testingT: t,
		assert:   assert.New(t),
		require:  require.New(t),
	}
}
