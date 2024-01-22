GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test
GOGET := $(GO) get

APP_NAME := your-app-name

VERSION := 1.0.0
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

SRC_DIR := src
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin

SRC := $(wildcard $(SRC_DIR)/*.go)

all: build

build: $(BIN_DIR)/$(APP_NAME)

$(BIN_DIR)/$(APP_NAME): $(SRC)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	$(GOBUILD) $(LDFLAGS) -o $@

clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BIN_DIR)

test:
	@echo "Running tests..."
	@$(GOTEST) ./...

run: build
	@echo "Running $(APP_NAME)..."
	@./$(BIN_DIR)/$(APP_NAME)

.PHONY: all build clean test run
