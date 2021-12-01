PATH := $(shell go env GOPATH)/bin:$(PATH)

install:                  ## Install programs.
	go install -v ./...

test:                     ## Run tests.
	go test -v -race ./...

bench:                    ## Run benchmarks.
	go test -v -race -run=XXX -bench=. ./...

tag: export VERSION ?= $(shell cat VERSION)
tag:
	test -n "$(VERSION)" || error "set VERSION environment variable"
	git tag -d v$(VERSION) || true
	git tag v$(VERSION) -m "Kata 形🤺 $(VERSION)"
	git push origin v$(VERSION)

release: export GITHUB_TOKEN ?= $(shell cat github.token)
release: export GPG_FINGERPRINT ?= $(shell cat gpg.fingerprint)
release: tag              ## Create new release. Requires GITHUB_TOKEN to be set.
	@# In order to release to GitHub,
	@# you'll need to export a GITHUB_TOKEN environment variable,
	@# which should contain a valid GitHub token with the repo scope.
	@# It will be used to deploy releases to your GitHub repository.
	@# You can create a new github token here: https://github.com/settings/tokens/new
	@# Run `VERSION=1.1.2 GITHUB_TOKEN=secret_token make release`.
	test -n "$(GITHUB_TOKEN)"    || error "set GITHUB_TOKEN environment variable"
	test -n "$(GPG_FINGERPRINT)" || error "set GPG_FINGERPRINT environment variable"
	go install github.com/goreleaser/goreleaser@latest
	#goreleaser init
	goreleaser check
	rm -rf dist
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
