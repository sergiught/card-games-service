package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/features/scenario"
	"github.com/sergiught/card-games-service/internal/config"
	"github.com/sergiught/card-games-service/internal/database"
)

const (
	testSuiteName   = "card-games"
	testSuiteFormat = "pretty"
	testSuiteDIR    = "./../../features"
)

func TestFeatures(t *testing.T) {
	configuration, err := config.LoadFromEnv()
	require.NoError(t, err)

	db, err := database.Connect(configuration.Database)
	require.NoError(t, err)

	deckCtx := scenario.NewDeckContext(configuration, db)

	suite := godog.TestSuite{
		Name: testSuiteName,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			scenario.RegisterCreateDeckSteps(ctx, deckCtx)
			scenario.RegisterOpenDeckSteps(ctx, deckCtx)
		},
		Options: &godog.Options{
			Format:   testSuiteFormat,
			Paths:    []string{testSuiteDIR},
			TestingT: t,
		},
	}

	exitCode := suite.Run()
	assert.Zero(t, exitCode, "failure encountered while running feature tests")
}
