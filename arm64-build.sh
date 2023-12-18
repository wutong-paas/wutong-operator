#! /bin/bash

WUTONG_REGISTRY=${WUTONG_REGISTRY:-'swr.cn-southwest-2.myhuaweicloud.com/wutong'}
IMAGE_NAME=wutong-operator
VERSION=v1.9.0-arm64

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o ./bin/arm64/manager main.go 

docker build . -t ${WUTONG_REGISTRY}/${IMAGE_NAME}:${VERSION} -f Dockerfile.local
docker push ${WUTONG_REGISTRY}/${IMAGE_NAME}:${VERSION}

rm -rf ./bin/*