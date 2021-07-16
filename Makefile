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
	go build -o poker

run:
	go run main.go deal

test:
	go test ./cmd/dealer/...

#can use a different location for a dockerfile by -f ./<path to dockerfile + name of dockerfile>
#TODO - MAKE THIS A SERVER THAT CAN BE INTERFACED VIA COMMAND LINE OR REST/GRPC CALLS
docker-build:
	docker build -f ./Dockerfile -t poker:$(TAG) .

docker-build-latest:
	docker build -t poker .

docker-run-latest:
	docker run --name api -d -p 4050:4050 poker

docker-run:
	docker run --name api -d -p 4050:4050 poker:$(TAG)

br: docker-build docker-run
