#! /bin/bash

export NAMESPACE=wutong-operator
export VERSION=v1.4.0

go mod download

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o bin/amd64/manager main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build -a -o bin/arm64/manager main.go

docker buildx use operatorbuilder || docker buildx create --use --name operatorbuilder
docker buildx build --platform linux/amd64,linux/arm64 --push -t swr.cn-southwest-2.myhuaweicloud.com/wutong/${NAMESPACE}:${VERSION} -f Dockerfile.local . 
# docker buildx rm operatorbuilder
