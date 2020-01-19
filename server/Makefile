GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

OUTPUT_DIR=build
BINARY_NAME=kwmassetsliabilities
TARGET=$(OUTPUT_DIR)/$(BINARY_NAME)

all: build test
.PHONY: build
build:
	$(GOBUILD) -o $(TARGET) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(TARGET)
run:
	$(GOBUILD) -o $(TARGET) -v
	./$(TARGET) server start
deps:
	$(GOGET)
