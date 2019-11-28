.PHONY: build vendor

all: clean build

build:
	@go mod vendor
	@go build -mod=vendor -o app main.go
	
clean:
	@rm -f app
	@go clean


