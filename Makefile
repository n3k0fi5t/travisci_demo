MODULE = $(shell go list -m)
LDFLAGS :=
EXEC := app

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: clean ## build the app binary
	CGO_ENABLED=0 go build ${LDFLAGS} -o ${EXEC} .

.PHONY: test
test: ## run unit tests
	go test -cover -covermode=count

.PHONY: run
run: build ## run the application
	./${EXEC}

.PHONY: clean
clean: ## remove temporary files
	rm -rf  ${EXEC}