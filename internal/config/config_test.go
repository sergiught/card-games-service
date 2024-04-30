package config_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sergiught/card-games-service/internal/config"
	"github.com/sergiught/card-games-service/internal/server"
)

func TestMain(m *testing.M) {
	// Save current environment and then clear it as
	// our config tests need a clean environment.
	originalEnvironment := os.Environ()
	os.Clearenv()

	// Run all the tests.
	exitCode := m.Run()

	// Restore original environment.
	for _, env := range originalEnvironment {
		environmentPair := strings.Split(env, "=")
		key := environmentPair[0]
		value := environmentPair[1]
		_ = os.Setenv(key, value)
	}

	// Exit with the correct code.
	os.Exit(exitCode)
}

func TestLoadFromEnv(t *testing.T) {
	testCases := []struct {
		name           string
		givenConfig    map[string]string
		expectedConfig *config.Specification
		expectedError  string
	}{
		{
			name: "it successfully loads all the given env vars",
			givenConfig: map[string]string{
				"ENVIRONMENT":             "production",
				"SERVER_ADDRESS":          "0.0.0.0:8080",
				"SERVER_READ_TIMEOUT":     "11s",
				"SERVER_WRITE_TIMEOUT":    "3s",
				"SERVER_SHUTDOWN_TIMEOUT": "9s",
			},
			expectedConfig: &config.Specification{
				Environment: "production",
				Server: server.Config{
					Address:         "0.0.0.0:8080",
					ReadTimeout:     time.Second * 11,
					WriteTimeout:    time.Second * 3,
					ShutdownTimeout: time.Second * 9,
				},
			},
		},
		{
			name:        "it successfully loads defaults if no env vars set",
			givenConfig: map[string]string{},
			expectedConfig: &config.Specification{
				Environment: "development",
				Server: server.Config{
					Address:         "0.0.0.0:8000",
					ReadTimeout:     time.Second * 5,
					WriteTimeout:    time.Second * 5,
					ShutdownTimeout: time.Second * 5,
				},
			},
		},
		{
			name: "it fails to load env vars if encountering a wrong type",
			givenConfig: map[string]string{
				"SERVER_READ_TIMEOUT": "tooHigh",
			},
			expectedError: "envconfig.Process: assigning SERVER_READ_TIMEOUT to ReadTimeout: converting 'tooHigh' to type time.Duration. details: time: invalid duration \"tooHigh\"",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for key, value := range testCase.givenConfig {
				// Do not run these in parallel, because
				// t.Setenv affects the whole process.
				t.Setenv(key, value)
			}

			actualConfig, err := config.LoadFromEnv()

			if testCase.expectedError != "" {
				require.EqualError(t, err, testCase.expectedError)
				return
			}

			assert.Equal(t, testCase.expectedConfig, actualConfig)
		})
	}
}
