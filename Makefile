PATH := $(shell go env GOPATH)/bin:$(PATH)

install:                  ## Install programs.
	go install -v ./...

test:                     ## Run tests.
	go test -v -race ./...

bench:                    ## Run benchmarks.
	go test -run=XXX -bench=. ./...

fmt:	                  ## Format code.
	go fmt ./...

lint:                     ## Lint code.
	golangci-lint run --enable-all

doc:                      ## Open godoc.
	pkill godoc; godoc &
	open "http://localhost:6060/pkg/$$(go list -m)"

help: Makefile            ## Display this help message.
	@echo "Use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | sort | awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

.DEFAULT_GOAL := help
.PHONY: $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))
