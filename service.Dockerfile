FROM golang:1.18rc1-alpine3.15

WORKDIR /build
COPY ./grpc ./grpc
COPY ./service ./service

RUN \
    --mount=type=cache,target=/go/pkg/mod \
    apk update && apk upgrade && apk add git upx gcc musl-dev &&\
    cd ./service &&\
    go test ./... && go build -o ../port-service -ldflags="-s -w" && upx ../port-service && upx -t ../port-service

FROM alpine:3.15

COPY --from=0 /build/port-service /usr/local/bin/port-service

RUN apk update && apk upgrade

EXPOSE 59001
EXPOSE 59002
ENTRYPOINT ["/usr/local/bin/port-service"]