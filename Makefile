all: build

build:
	@go build -o build/rest-api-task cmd/rest-api-task/main.go

install:
	@GOPRIVATE=github.com/bleenco/* go mod download

.PHONY: build install
