#! /bin/bash

WUTONG_REGISTRY=${WUTONG_REGISTRY:-'swr.cn-southwest-2.myhuaweicloud.com/wutong'}
IMAGE_NAME=wutong-operator
VERSION=v1.15.0-amd64

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/amd64/manager main.go

docker build . -t ${WUTONG_REGISTRY}/${IMAGE_NAME}:${VERSION} -f Dockerfile.local
docker push ${WUTONG_REGISTRY}/${IMAGE_NAME}:${VERSION}

rm -rf ./bin/*