APP=${shell basename $(shell git remote get-url origin)}
REGISTRY=evsmoker
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
TARGETARCH=arm64 # amd64

format:
	gofmt -s -w ./

lint:
	golint

test:
	@echo "Running tests..."
	go test -v

get:
	@echo "Getting dependencies..."
	go get

build: format get
	@echo "Building production version..."
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o telebot-go -ldflags "-X=github.com/ev-smoke/telebot-go/cmd.appVersion=${VERSION}"

image:
	@echo "Building Docker image..."
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

push:
	@echo "Pushing Docker image..."
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARCH}

clean:
	@echo "Cleaning build artifacts..."
	rm -rf telebot-go

help:
	@echo "Make Targets:"
	@echo "make help			- Show this help message"
	@echo "make tests			- Run tests"
	@echo "make image			- make app docker image"
	@echo "make push			- push docker to registry"
	@echo "make clean			- clean autogenerated files + clean"
	@echo "make plarform 			- make app for specific platform"