# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=cryptotracker
BINARY_UNIX=$(BINARY_NAME)_unix

all: build run

build:
	@echo Building $(BINARY_NAME) project:
	$(GOBUILD) -o $(BINARY_NAME) -v
	@echo Project built!!

clean:
	@echo Cleaning $(BINARY_NAME) project files:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	@echo Cleaning Finished!!

run:
	@$(GOBUILD) -o $(BINARY_NAME) ./...
	@./$(BINARY_NAME)

deps:
	@echo Gathering dependencies:
	$(GOGET) -u -v github.com/miguelmota/go-coinmarketcap gopkg.in/urfave/cli.v2
	@echo Dependencies up to date!
