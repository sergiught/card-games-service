package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"

	"github.com/sergiught/card-games-service/features/scenario"
)

const (
	testSuiteName   = "card-games"
	testSuiteFormat = "pretty"
	testSuiteDIR    = "./../../features"
)

func TestFeatures(t *testing.T) {
	ctx := scenario.NewContext(t)

	suite := godog.TestSuite{
		Name:                testSuiteName,
		ScenarioInitializer: scenario.Initialize(ctx),
		Options: &godog.Options{
			Format:   testSuiteFormat,
			Paths:    []string{testSuiteDIR},
			TestingT: t,
		},
	}

	exitCode := suite.Run()
	assert.Zero(t, exitCode, "failure encountered while running feature tests")
}
