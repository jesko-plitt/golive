TESTPACKAGES = `go list ./... | grep -v test`

.PHONY: help
help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: test
test: ## Test the application
	$(go-docker-run-base) go test `go list ./... | grep -v test` -count=1 -timeout 15s

.PHONY: cover
cover: ## Run test coverage
	go test $(TESTPACKAGES) -coverprofile=coverage.out -covermode=atomic -count=1 -timeout 25s
	go tool cover -html=coverage.out
	go tool cover -func coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'
	rm coverage.out
