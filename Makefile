NAME ?= diego-exit-code-catcher
OUTPUT = ./bin/$(NAME)
CF_OUTPUT = ./bin/cf/$(NAME)
GO_SOURCES = $(shell find . -type f -name '*.go')
VERSION ?= 0.0.0

.PHONY: all
all: build

.PHONY: clean
clean:
	@rm -rf bin/

$(OUTPUT): $(GO_SOURCES)
	@echo "Building $(VERSION)"
	@mkdir -p ./bin/cf
	go build -o $(OUTPUT) .
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -o $(CF_OUTPUT) .

.PHONY: build
build: $(OUTPUT) ## Build the main binary

.PHONY: tidy
tidy: ## Remove unused dependencies
	go mod tidy
