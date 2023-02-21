FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/kbrprime-be
ENV GOSUMDB=off
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /go/bin/kbrprime-be main.go
FROM alpine:3.12
RUN apk add --no-cache tzdata ca-certificates libc6-compat
COPY --from=builder /go/bin/kbrprime-be /go/bin/kbrprime-be
COPY --from=builder /go/src/kbrprime-be/.env /go/src/kbrprime-be/.env

ENTRYPOINT ["/go/bin/kbrprime-be"]