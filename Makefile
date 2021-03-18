SHELL := /bin/bash

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP := $(shell date "+%Y%m%d%H%M%S")

$(eval TAG=$(CURRENT_BRANCH)_$(TIMESTAMP))

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
	docker build -f ./Dockerfile -t poker-app:$(TAG) .

docker-run:
	docker run --name api -d -p 4050:4050 poker-app:$(TAG)

br: docker-build docker-run

docker-push:
	docker image push poker-app:$(TAG)