package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/features/scenario"
	"github.com/sergiught/card-games-service/internal/config"
)

const (
	testSuiteName   = "card-games"
	testSuiteFormat = "pretty"
	testSuiteDIR    = "./../../features"
)

func TestFeatures(t *testing.T) {
	configuration, err := config.LoadFromEnv()
	require.NoError(t, err)

	deckCtx := scenario.NewDeckContext(configuration)

	suite := godog.TestSuite{
		Name:                testSuiteName,
		ScenarioInitializer: scenario.Initialize(deckCtx),
		Options: &godog.Options{
			Format:   testSuiteFormat,
			Paths:    []string{testSuiteDIR},
			TestingT: t,
		},
	}

	exitCode := suite.Run()
	assert.Zero(t, exitCode, "failure encountered while running feature tests")
}
