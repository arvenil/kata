PATH := $(shell go env GOPATH)/bin:$(PATH)

install:                  ## Install programs.
	go install -v ./...

test:                     ## Run tests.
	go test -v -race ./...

bench:                    ## Run benchmarks.
	go test -v -race -run=XXX -bench=. ./...

release: export GPG_FINGERPRINT:=284E7B63A3840723B4FC8354A31DBFF63363B349
release:                  ## Create new release. Requires GITHUB_TOKEN to be set.
						  ## In order to release to GitHub,
						  ## you'll need to export a GITHUB_TOKEN environment variable,
						  ## which should contain a valid GitHub token with the repo scope.
						  ## It will be used to deploy releases to your GitHub repository.
						  ## You can create a new github token here: https://github.com/settings/tokens/new
						  ## Run `VERSION=1.1.2 GITHUB_TOKEN=secret_token make release`.
	go install github.com/goreleaser/goreleaser@latest
	#goreleaser init
	goreleaser check
	rm -rf dist
	git tag "v$(VERSION)" -m "Kata 形🤺 $(VERSION)" || true
	goreleaser release

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
