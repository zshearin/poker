SHELL := /bin/bash

#-include $(shell curl -sSL -o .build-harness "https://git.io/build-harness"; echo .build-harness)

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP := $(shell date "+%Y%m%d%H%M%S")

#$(eval TAG=$(CURRENT_BRANCH)_$(TIMESTAMP))
$(eval TAG=$(CURRENT_BRANCH))

build-linux:
	GOOS=linux GOARCH=amd64 go build ./cmd/main.go && mv main ./bin/application

tag:
	@echo $(TAG)

build:
	go build ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./poker/...

#can use a different location for a dockerfile by -f ./<path to dockerfile + name of dockerfile>
docker-build:
	docker build -f ./Dockerfile -t poker:$(TAG) .

docker-build-latest:
	docker build -t poker .

docker-run-latest:
	docker run --name api -d -p 4050:4050 poker

docker-run:
	docker run --name api -d -p 4050:4050 poker:$(TAG)

br: docker-build docker-run


proto:
	cd cmd
	protoc --go_out=plugins=grpc:proto poker.proto
