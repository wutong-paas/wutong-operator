FROM swr.cn-southwest-2.myhuaweicloud.com/wutong/alpine:3.15
RUN mkdir /app \
    && apk add --update apache2-utils \
    && rm -rf /var/cache/apk/*
ENV TZ=Asia/Shanghai
WORKDIR /
COPY ./bin/app ./manager

CMD ["./manager"]