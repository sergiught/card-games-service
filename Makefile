#-----------------------------------------------------------------------------------------------------------------------
# Variables (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
#-----------------------------------------------------------------------------------------------------------------------
BINARY_NAME = card-games-service
BUILD_DIR ?= $(CURDIR)/dist
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

$(GO_BIN)/CompileDaemon:
	@echo "==> Installing CompileDaemon within ${GO_BIN}"
	@go install -v github.com/githubnemo/CompileDaemon@50a8debecc13686ba29356b28d30c33555e662f6 # v1.4.0

build:
	@echo "==> Building the Card Games Service binary within ${BUILD_DIR}/${BINARY_NAME}"
	@go build -v -o "${BUILD_DIR}/${BINARY_NAME}" "$(CURDIR)/cmd/api/main.go"

install: $(GO_BIN)/CompileDaemon build
	@echo "==> Installing the Card Games Service binary within ${GO_BIN}"
	@mv "${BUILD_DIR}/${BINARY_NAME}" "${GO_BIN}"

lint: $(GO_BIN)/golangci-lint
	@echo "==> Running golangci-lint over project"
	@golangci-lint run -v --fix -c .golangci.yml ./...

lint-commits: $(GO_BIN)/commitlint
	@commitlint lint

test-unit:
	@echo "==> Running unit tests"
	@go test -v -race -count=1 ./...

test-features:
	@echo "==> Running feature tests"
	@docker-compose exec card-games-service go test -v -race -count=1 -tags=features -timeout=10m -run TestFeatures/"$(FILTER)" ./cmd/api/main_test.go

dev-migrate-create: $(GO_BIN)/migrate
	@echo "==> Creating database migration files"
	@migrate create -ext "sql" -dir ${DB_MIGRATIONS_PATH} ${DB_MIGRATIONS_FILE_NAME}

dev-migrate-up: $(GO_BIN)/migrate
	@echo "==> Running migrations against the local postgres development database"
	@migrate -database "${DB_DSN}" -path "${DB_MIGRATIONS_PATH}" up

dev-migrate-down: $(GO_BIN)/migrate
	@echo "==> Running migrations against the local postgres development database"
	@migrate -database "${DB_DSN}" -path "${DB_MIGRATIONS_PATH}" down

dev-env:
	@echo "==> Generating .env file"
	@cp .env.example .env

dev-run: dev-env
	@echo "==> Starting development environment"
	@docker compose up -d --build
	@sleep 2
	@$(MAKE) dev-migrate-up
	@docker-compose logs -f card-games-service
