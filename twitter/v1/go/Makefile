OUT_DIR="./src"
PROJECT_ID=kooblabs
IMAGE_NAME=lcdemos-twitter-v1
CLUSTER_NAME=lcdemos-twitter-v1
GOPATH := $(shell go env GOPATH)

all: proto

proto:
	mkdir -p gen
	PATH="$(PATH):$(GOPATH)/bin" protoc -I ../../protos/ ../../protos/*.proto --go_out=plugins=grpc:./gen --go_opt=paths=source_relative


run:
	PATH="$(PATH):$(GOPATH)/bin" go run server.go
