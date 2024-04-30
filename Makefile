#-----------------------------------------------------------------------------------------------------------------------
# Variables (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
#-----------------------------------------------------------------------------------------------------------------------
GO_BIN ?= $(shell go env GOPATH)/bin

DB_MIGRATIONS_PATH ?= "$(CURDIR)/infrastructure/migrations"
DB_MIGRATIONS_FILE_NAME ?= "new_migration"
DB_DSN ?= "postgresql://dealer:password@localhost:5432/card_decks?sslmode=disable"

#-----------------------------------------------------------------------------------------------------------------------
# Rules (https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html#Rule-Introduction)
#-----------------------------------------------------------------------------------------------------------------------
$(GO_BIN)/golangci-lint:
	@echo "==> Installing golangci-lint within ${GO_BIN}"
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@77a8601aa372eaab3ad2d7fa1dffa71f28005393 # v1.57.2

$(GO_BIN)/commitlint:
	@echo "==> Installing commitlint within ${GO_BIN}"
	@go install -v github.com/conventionalcommit/commitlint@e9a606ce7074ac884ea091765be1651be18356d4 # v0.10.1

$(GO_BIN)/migrate:
	@echo "==> Installing migrate within ${GO_BIN}"
	@go install -v -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@cd17c5a808d1889d17d73a68355c18d1dea6c49d # v4.17.0

lint: $(GO_BIN)/golangci-lint
	@echo "==> Running golangci-lint over project"
	@golangci-lint run -v --fix -c .golangci.yml ./...

lint-commits: $(GO_BIN)/commitlint
	@commitlint lint

test-features:
	@echo "==> Running feature tests"
	@go test -v -race -count=1 -timeout=10m -run TestFeatures/"$(FILTER)" ./cmd/api/main_test.go

dev-migrate-create: $(GO_BIN)/migrate
	@echo "==> Creating database migration files"
	@migrate create -ext "sql" -dir ${DB_MIGRATIONS_PATH} ${DB_MIGRATIONS_FILE_NAME}

dev-migrate-up: $(GO_BIN)/migrate
	@echo "==> Running migrations against the local postgres development database"
	@migrate -database "${DB_DSN}" -path "${DB_MIGRATIONS_PATH}" up

dev-migrate-down: $(GO_BIN)/migrate
	@echo "==> Running migrations against the local postgres development database"
	@migrate -database "${DB_DSN}" -path "${DB_MIGRATIONS_PATH}" down
