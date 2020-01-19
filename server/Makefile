GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u

OUTPUT_DIR=build
BINARY_NAME=kwmassetsliabilities
TARGET=$(OUTPUT_DIR)/$(BINARY_NAME)

ifneq ("$(wildcard ./.env)","")
	include ./.env
	export $(shell sed 's/=.*//' ./.env)
endif

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

build-migrations:
	go build -o ./migrations/goose ./cmd/goose/*.go 
	chmod +x ./migrations/goose

db-new-migration:
	goose -dir migrations/ create ${name}

db-migrate: build-migrations
	./migrations/goose -dir migrations postgres "host=${AL_DB_HOST} port=${AL_DB_PORT} user=${AL_DB_USER}  password=${AL_DB_PASSWORD} dbname=${AL_DB_NAME} sslmode=${AL_DB_SSL_MODE}" up

db-rollback: build-migrations
	./migrations/goose -dir migrations postgres "host=${AL_DB_HOST} port=${AL_DB_PORT} user=${AL_DB_USER}  password=${AL_DB_PASSWORD} dbname=${AL_DB_NAME} sslmode=${AL_DB_SSL_MODE}" down

