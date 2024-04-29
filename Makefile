#-----------------------------------------------------------------------------------------------------------------------
# Variables (https://www.gnu.org/software/make/manual/html_node/Using-Variables.html#Using-Variables)
#-----------------------------------------------------------------------------------------------------------------------
GO_BIN ?= $(shell go env GOPATH)/bin

#-----------------------------------------------------------------------------------------------------------------------
# Rules (https://www.gnu.org/software/make/manual/html_node/Rule-Introduction.html#Rule-Introduction)
#-----------------------------------------------------------------------------------------------------------------------
$(GO_BIN)/commitlint:
	@echo "==> Installing commitlint within ${GO_BIN}"
	@go install -v github.com/conventionalcommit/commitlint@e9a606ce7074ac884ea091765be1651be18356d4 # v0.10.1

lint-commits: $(GO_BIN)/commitlint
	@commitlint lint
