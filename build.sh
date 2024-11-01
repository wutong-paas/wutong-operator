#! /bin/bash

IMAGE_REGISTRY=${WUTONG_REGISTRY:-'swr.cn-southwest-2.myhuaweicloud.com/wutong'}
IMAGE_NAME=wutong-operator
VERSION=v1.15.1-alpha.1

go mod download

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/amd64/manager main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o bin/arm64/manager main.go

docker buildx use swrbuilder || docker buildx create --use --name swrbuilder --driver docker-container --driver-opt image=swr.cn-southwest-2.myhuaweicloud.com/wutong/buildkit:stable
docker buildx build --platform linux/amd64,linux/arm64 --push -t ${IMAGE_REGISTRY}/${IMAGE_NAME}:${VERSION} -f Dockerfile.local . 
# docker buildx rm swrbuilder

rm -rf ./bin/*
