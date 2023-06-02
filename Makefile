all: test build

.PHONY: test
test:  ## Run tests
	go test ./.../

.PHONY: build
build:  ## Build binary
	go build

.PHONY: fmt
fmt:  ## Format code
	go fmt ./.../

.PHONY: help
help:
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_\/-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
