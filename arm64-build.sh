#! /bin/bash

export NAMESPACE=wutong-operator
export VERSION=v1.4.0-arm64

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -o ./bin/arm64/manager main.go 

docker build . -t swr.cn-southwest-2.myhuaweicloud.com/wutong/${NAMESPACE}:${VERSION} -f Dockerfile.local
docker push swr.cn-southwest-2.myhuaweicloud.com/wutong/${NAMESPACE}:${VERSION}

rm -rf ./bin/*