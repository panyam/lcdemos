OUT_DIR="./src"
PROJECT_ID=kooblabs
IMAGE_NAME=lcdemos-twitter-v1
CLUSTER_NAME=lcdemos-twitter-v1
GOPATH := $(shell go env GOPATH)
PROTOSRC=../protos

all: proto

proto:
	mkdir -p gen
	PATH="$(PATH):$(GOPATH)/bin" protoc -I $(PROTOSRC)/ $(PROTOSRC)/*.proto --go_out=plugins=grpc:./gen --go_opt=paths=source_relative
	PATH="$(PATH):$(GOPATH)/bin" protoc -I $(PROTOSRC)/ $(PROTOSRC)/*.proto --swagger_out=logtostderr=true:./gen
	PATH="$(PATH):$(GOPATH)/bin" protoc -I $(PROTOSRC)/ $(PROTOSRC)/*.proto --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen
	go install

docker: proto
	PATH=~/.pyenv/versions/2.7.17/bin/:$(PATH) docker build -f app/Dockerfile -t gcr.io/${PROJECT_ID}/${IMAGE_NAME}:latest .

push:
	docker push gcr.io/${PROJECT_ID}/${IMAGE_NAME}:latest

ensurecluster:
	kubectl create deployment ${CLUSTER_NAME} --image=gcr.io/${PROJECT_ID}/lcdemos-twitter-v1:latest
	kubectl scale deployment ${CLUSTER_NAME} --replicas=3
