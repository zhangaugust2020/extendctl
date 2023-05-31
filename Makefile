# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=extendctl
BINARY_UNIX=$(BINARY_NAME)_unix
DATETIME=$(shell date.exe "+%Y%m%d%H%M%S")

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
#deps:
#	$(GOGET) github.com/markbates/goth
#	$(GOGET) github.com/markbates/pop

# Cross compilation
linux:
	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	$(GOBUILD) -o $(BINARY_NAME) .
#docker-build:
#	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v