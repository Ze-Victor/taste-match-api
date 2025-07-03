.PHONY: default run build test lint clean

APP_NAME=taste-match-api

default: run

run:
	@go run main.go

dependencies:
	@go mod tidy

build:
	@go build -o $(APP_NAME) main.go
