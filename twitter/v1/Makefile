OUT_DIR="./"

all: proto docker

proto:
	cd src &&	\
	 python -m grpc_tools.protoc -I./protos  --python_out=$(OUT_DIR) --grpc_python_out=$(OUT_DIR) ./protos/*.proto &&	\
	 python -m grpc_tools.protoc -I./protos  --python_out=./ --grpc_python_out=./ ./protos/google/api/*.proto 

docker:
	docker build -f app/Dockerfile -t lcdemos-twitter-v1:latest .
