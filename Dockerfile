FROM golang:1.10.3-alpine3.8 as builder

RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && update-ca-certificates \
    && apk add git

WORKDIR /go/src/github.com/mtrojanowski-kainos/keyvault-flex-demo
COPY . .

RUN go get -d ./...
RUN go build -o /demo main.go

FROM alpine:3.8
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/* \
    && update-ca-certificates

EXPOSE 8080
ENTRYPOINT [ "/demo" ]
COPY --from=builder /demo /
