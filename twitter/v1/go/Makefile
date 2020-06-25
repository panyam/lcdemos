OUT_DIR="./src"
PROJECT_ID=kooblabs
IMAGE_NAME=lcdemos-twitter-v1
CLUSTER_NAME=lcdemos-twitter-v1
GOPATH := $(shell go env GOPATH)

all: proto

proto:
	mkdir -p gen
	PATH="$(PATH):$(GOPATH)/bin" protoc -I ../../protos/ ../../protos/*.proto --go_out=plugins=grpc:./gen --go_opt=paths=source_relative
	PATH="$(PATH):$(GOPATH)/bin" protoc -I ../../protos/ ../../protos/*.proto --swagger_out=logtostderr=true:./gen
	PATH="$(PATH):$(GOPATH)/bin" protoc -I ../../protos/ ../../protos/*.proto --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen
	go install

