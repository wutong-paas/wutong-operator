#! /bin/bash

export IMAGE_REGISTRY=swr.cn-southwest-2.myhuaweicloud.com/wutong
export IMAGE_NAME=wutong-operator
export VERSION=v1.4.0

go mod download

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/amd64/manager main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o bin/arm64/manager main.go

docker buildx use gobuilder || docker buildx create --use --name gobuilder
docker buildx build --platform linux/amd64,linux/arm64 --push -t ${IMAGE_REGISTRY}/${IMAGE_NAME}:${VERSION} -f Dockerfile.local . 
# docker buildx rm gobuilder

rm -rf ./bin/*