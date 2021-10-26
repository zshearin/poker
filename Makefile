SHELL := /bin/bash

#-include $(shell curl -sSL -o .build-harness "https://git.io/build-harness"; echo .build-harness)

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP := $(shell date "+%Y%m%d%H%M%S")

#$(eval TAG=$(CURRENT_BRANCH)_$(TIMESTAMP))
$(eval TAG=$(CURRENT_BRANCH))

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1alpha1/poker.proto

build-linux:
	GOOS=linux GOARCH=amd64 go build ./cmd/main.go && mv main ./bin/application

tag:
	@echo $(TAG)

build: check-protoc proto build-bin

check-protoc:
	@echo "Checking if protoc command line tool is installed"
	protoc --version

build-bin:
	go build -o poker

run:
	go run main.go deal

clean:
	rm poker

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
