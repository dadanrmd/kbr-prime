FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/kbrprime-be
RUN pwd
COPY . .
ENV GOSUMDB=off
COPY go.mod .
COPY params/.env /params/.env
COPY go.sum .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main .

FROM golang:1.18-alpine

RUN apk add --no-cache tzdata ca-certificates libc6-compat

COPY --from=builder $PWD /usr/local/go/bin/kbrprime-be
COPY --from=builder $PWD/params/.env /usr/local/go/src/kbrprime-be/params/.env

ENTRYPOINT ["/go/bin/kbrprime-be"]

