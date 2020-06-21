OUT_DIR="./"
PROJECT_ID=kooblabs
IMAGE_NAME=lcdemos-twitter-v1
CLUSTER_NAME=lcdemos-twitter-v1

all: proto docker

proto:
	 python -m grpc_tools.protoc -I../protos  -I../../google --python_out=./src/$(OUT_DIR) --grpc_python_out=./src/$(OUT_DIR) ../protos/*.proto
	 python -m grpc_tools.protoc -I../protos --python_out=./src/ --grpc_python_out=./src ../protos/google/api/*.proto 

docker:
	docker build -f app/Dockerfile -t gcr.io/${PROJECT_ID}/${IMAGE_NAME}:latest .

push:
	docker push gcr.io/${PROJECT_ID}/${IMAGE_NAME}:latest

ensurecluster:
	kubectl create deployment ${CLUSTER_NAME} --image=gcr.io/${PROJECT_ID}/lcdemos-twitter-v1:latest
	kubectl scale deployment ${CLUSTER_NAME} --replicas=3
