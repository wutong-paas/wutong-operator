#! /bin/bash

export NAMESPACE=wutong-operator
export VERSION=v1.4.0
docker buildx create --use --name operatorbuilder || docker buildx use operatorbuilder
docker buildx build --platform linux/amd64,linux/arm64 --push -t swr.cn-southwest-2.myhuaweicloud.com/wutong/${NAMESPACE}:${VERSION} -f Dockerfile.multiarch . 
# docker buildx rm operatorbuilder