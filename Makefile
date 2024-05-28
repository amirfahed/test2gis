APP_NAME = test2gis

SRC_DIR = ./cmd
TEST_DIR = ./internal/service

DEPS = github.com/go-chi/chi/v5

GOLANGCI_TAG ?= 1.59.0

# install dependency
.PHONY: deps
deps:
	go get -u $(DEPS)

# build app
.PHONY: build
build:
	go build -o $(APP_NAME) $(SRC_DIR)/main.go

# run app
.PHONY: run
run: build
	./$(APP_NAME)

# run test
.PHONY: test
test:
	go test -v $(TEST_DIR)

# clean
.PHONY: clean
clean:
	rm -f $(APP_NAME)

# install deps and build
.PHONY: setup
setup: deps build

# install golangci-lint
.PHONY: install-lint
install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)

# run linter
.PHONY: lint
lint: install-lint
	golangci-lint run