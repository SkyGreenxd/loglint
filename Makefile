PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

GOLANGCI_LINT = $(PROJECT_BIN)/golangci-lint
CUSTOM_GCL    = $(PROJECT_DIR)/custom-gcl

.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(GOLANGCI_LINT) ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v2.8.0

.PHONY: all
all: .install-linter ## Build and run custom golangci-lint with loglint
	$(GOLANGCI_LINT) custom -v
	$(CUSTOM_GCL) run ./...

.PHONY: test
test: ## Run all tests
	go test -v -race -count=1 ./analyzer/...

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(PROJECT_BIN)
	rm -f $(CUSTOM_GCL)

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'