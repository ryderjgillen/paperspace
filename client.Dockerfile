FROM golang:1.18rc1-alpine3.15

WORKDIR /build
COPY ./grpc ./grpc
COPY ./client ./client

RUN \
    --mount=type=cache,target=/go/pkg/mod \
    apk update && apk upgrade && apk add git upx &&\
    cd ./client &&\
    go build -o ../port-service-client -ldflags="-s -w" && upx ../port-service-client && upx -t ../port-service-client

FROM alpine:3.15

COPY --from=0 /build/port-service-client /usr/local/bin/port-service-client

RUN apk update && apk upgrade

ENTRYPOINT ["/usr/local/bin/port-service-client"]