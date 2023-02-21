FROM golang:1.16.2-alpine3.12 AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/kbrprime-be
RUN pwd
COPY . .
ENV GOSUMDB=off
COPY go.mod .
COPY params/.env .
COPY go.sum .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /go/bin/kbrprime-be pkg/main.go

FROM alpine:3.12

RUN apk add --no-cache tzdata ca-certificates libc6-compat

COPY --from=builder /go/bin/kbrprime-be /go/bin/kbrprime-be
COPY --from=builder /go/src/kbrprime-be/.env /go/src/kbrprime-be/.env

ENTRYPOINT ["/go/bin/kbrprime-be"]

