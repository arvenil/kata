install:                  ## Install programs.
	go install -v ./...

test:                     ## Run tests.
	go test -v -race ./...

fmt:	                  ## Format code.
	go fmt ./...

help: Makefile            ## Display this help message.
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		sort | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: install test fmt help
