FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*
COPY ./test /bin

ENTRYPOINT ["/bin/test"]