# Build the manager binary
FROM golang:1.21 as builder
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
ENV GOPROXY=https://goproxy.cn,direct
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY util/ util/

# Build
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GO111MODULE=on go build -a -o manager main.go


FROM swr.cn-southwest-2.myhuaweicloud.com/wutong/alpine:3.15
RUN mkdir /app \
    && apk add --update apache2-utils \
    && rm -rf /var/cache/apk/*
ENV TZ=Asia/Shanghai
WORKDIR /
COPY --from=builder /workspace/manager .

CMD ["/manager"]
