SHELL := /bin/bash

#-include $(shell curl -sSL -o .build-harness "https://git.io/build-harness"; echo .build-harness)

CURRENT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
TIMESTAMP := $(shell date "+%Y%m%d%H%M%S")

#$(eval TAG=$(CURRENT_BRANCH)_$(TIMESTAMP))
$(eval TAG=$(CURRENT_BRANCH))

<<<<<<< HEAD
=======
proto-new:
	mkdir -p backend/api/generated
	protoc -I. -I $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/ --go_out=plugins=grpc:backend/api/generated/. --grpc-gateway_out=logtostderr=true:backend/api/generated/. backend/api/v1alpha1/poker.proto
	mv backend/api/generated/github.com/zshearin/poker/api/v1alpha1/**.go backend/api/v1alpha1/
	rm -r backend/api/generated/

>>>>>>> e870b7b (make gRPC work and have http/gRPC server)
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/v1alpha1/poker.proto

build-linux:
	GOOS=linux GOARCH=amd64 go build ./backend/cmd/main.go && mv main ./bin/application

tag:
	@echo $(TAG)

<<<<<<< HEAD
build: check-protoc proto build-bin
=======
build: check-protoc proto-new build-bin
>>>>>>> e870b7b (make gRPC work and have http/gRPC server)

check-protoc:
	@echo "Checking if protoc command line tool is installed"
	protoc --version

build-bin:
<<<<<<< HEAD
	go build -o poker
=======
	cd backend; go build -o poker; cd ..
>>>>>>> e870b7b (make gRPC work and have http/gRPC server)

run:
	cd backend; go run main.go deal; cd ..

clean:
	cd backend; rm poker; cd ..

clean:
	rm poker

test:
	cd backend; go test ./cmd/dealer/...; cd ..

#can use a different location for a dockerfile by -f ./<path to dockerfile + name of dockerfile>
#TODO - MAKE THIS A SERVER THAT CAN BE INTERFACED VIA COMMAND LINE OR REST/GRPC CALLS
# docker-build:
# 	cd backend; docker build -f ./Dockerfile -t poker:$(TAG) .; cd ..

# docker-build-latest:
# 	cd backend; docker build -t poker .; cd ..

# docker-run-latest:
# 	cd backend; docker run --name api -d -p 4050:4050 poker; cd ..

# docker-run:
# 	cd backend; docker run --name api -d -p 4050:4050 poker:$(TAG); cd ..

# br: docker-build docker-run
